<?php
require './vendor/autoload.php';
use \Firebase\JWT\JWT;


/** 
 * 
 * 
 * 
*/
$LoginCallback = function($val) {
    header('Content-Type: application/json');
    //echo "CreateCallback! ".$val."\n";
    $dbHelper = new DBHelper("host=postgres port=5432 user=postgres password=password port=5432 dbname=switter");
    $rbody = json_decode( file_get_contents('php://input') ,true) ;
    $errno = json_last_error();
    if ( $errno == null ){
        if( $rbody["userEmail"] != null && $rbody["password"] != null ) {
            // user existing db check
            //$encodedPassword = password_hash($rbody["password"], PASSWORD_DEFAULT);
            if ( $dbHelper->checkUserExist($rbody["userEmail"]) ){
                if ($dbHelper->checkPassword($rbody["userEmail"] , $rbody["password"] )) {
                    $jwt = createJWT($rbody["userEmail"], $_ENV["JWT_SIGNING_KEY"], $_ENV["EXP_TIME"]);
                    $rt = createRT($jwt, $_ENV["RT_SIGNING_KEY"]);
                    $encodedRT = password_hash($rt, PASSWORD_DEFAULT);
                    $dbHelper->updateToken($rbody["userEmail"], $encodedRT );
                    $responce = array(
                        "jwt" => $jwt,
                        "rt" => $rt,
                    );
                    http_response_code(200);
                    print(json_encode($responce));
                }   else {
                    http_response_code(401);
                    print("PasswordNotMatch");    
                }
            } else {
                http_response_code(404);
                print("UserNotExist");
            }
        } else {
            http_response_code(401);
            print("AuthDataNotFull");
        }
    } else {
        http_response_code(400);
        print("NoData");
    }
    //print( $_ENV["SIGNING_KEY"]) ;
};
/** 
 * 
 * 
 * 
*/
$RegisterCallback = function($val) {
    header('Content-Type: application/json');
    //echo "CreateCallback! ".$val."\n";
    $dbHelper = new DBHelper("host=postgres port=5432 user=postgres password=password port=5432 dbname=switter");
    $rbody = json_decode( file_get_contents('php://input'),true ) ;
    $errno = json_last_error();
    if ( $errno == null ){
        if( $rbody["userEmail"] != null && $rbody["password"] != null && $rbody["userName"] != null ) {
            // user existing db check
            
            if (!$dbHelper->checkUserExist($rbody["userEmail"]  )) {
                $jwt = createJWT($rbody["userEmail"], $_ENV["JWT_SIGNING_KEY"], $_ENV["EXP_TIME"]);
                $rt = createRT($jwt, $_ENV["RT_SIGNING_KEY"]);
                $encodedRT = password_hash($rt, PASSWORD_DEFAULT);
                $encodedPassword = password_hash($rbody["password"], PASSWORD_DEFAULT);
                $dbHelper->createUser( $rbody["userName"],
                                        $rbody["userEmail"],
                                        $encodedPassword, 
                                        $encodedRT );
                $responce = array(
                    "jwt" => $jwt,
                    "rt" => $rt,
                );
                http_response_code(200);
                print(json_encode($responce));
            }   else {
                http_response_code(401);
                print("UserExist");    
            }
        } else {
            http_response_code(401);
            print("AuthDataNotFull");
        }
    } else {
        http_response_code(400);
        print("NoData");
    }
};
/** 
 * 
 * 
 * 
*/
$UpdateCallback = function($val) {
    //echo "EditCallback ".$val."!\n";
    header('Content-Type: application/json');
    $dbHelper = new DBHelper("host=postgres port=5432 user=postgres password=password port=5432 dbname=switter");
    $rbody = json_decode( file_get_contents('php://input'),true ) ;
    $errno = json_last_error();
    if ( $errno == null ){
        if( $rbody["jwt"] != null && $rbody["rt"] != null ){
            if (checkTokens( $rbody["jwt"], $rbody["rt"]) ) {
                http_response_code(401);
                print("JWTRTIncompatible");
                return;
            }
            //$decoded = JWT::decode( $rbody["jwt"], $_ENV["JWT_SIGNING_KEY"], array('HS256'));
            //$decoded_array = (array) $decoded;
            $userEmail = extractUserEmailFromJWT($rbody["jwt"]);

            if ( $userEmail != null ){
                if ($dbHelper->checkRT($userEmail, $rbody["rt"]) ){
                    $jwt = createJWT($userEmail, $_ENV["JWT_SIGNING_KEY"], $_ENV["EXP_TIME"]);
                    $rt = createRT($jwt, $_ENV["RT_SIGNING_KEY"]);
                    if ( $jwt != null && $rt != null ){
                        $encodedRT = password_hash($rt, PASSWORD_DEFAULT);
                        $dbHelper->updateToken($userEmail, $encodedRT );
                        $responce = array(
                            "jwt" => $jwt,
                            "rt" => $rt,
                        );
                        http_response_code(200);
                        print(json_encode($responce));
                    } else {
                        http_response_code(500);
                        print("SHeetHaPPenz");
                    }
                } else {
                    http_response_code(401);
                    print("InvalidRefreshToken");                
                };
            } else {
                http_response_code(401);
                print("IncorrectJWT");                
            };
        } else {
            http_response_code(401);
            print("DataNotFull");
        }
    } else {
        http_response_code(400);
        print("NoData");
    }
};
/** 
 * 
 * 
 * 
*/
$DeleteCallback = function($val) {
    //echo "DeleteCallback ".$val."!\n";
    header('Content-Type: application/json');
    $dbHelper = new DBHelper("host=postgres port=5432 user=postgres password=password port=5432 dbname=switter");
    $rbody = json_decode( file_get_contents('php://input'),true ) ;
    $errno = json_last_error();
    if ( $errno == null ){
        if( $rbody["jwt"] != null && $rbody["rt"] != null ){
            if (checkTokens( $rbody["jwt"], $rbody["rt"]) ) {
                http_response_code(401);
                print("JWTRTIncompatible");
            }
            $decoded = JWT::decode( $rbody["jwt"], $_ENV["JWT_SIGNING_KEY"], array('HS256'));
            $decoded_array = (array) $decoded;
            
            if ( $decoded_array["Email"] != null ){
                if ($dbHelper->checkRT($decoded_array["Email"], $rbody["rt"]) ){
                    $dbHelper->deleteToken($decoded_array["Email"]);
                    http_response_code(200);
                    print("logout succesfully");
                } else {
                    http_response_code(401);
                    print("InvalidRefreshToken");                
                };
            } else {
                http_response_code(401);
                print("IncorrectJWT");                
            };
        } else {
            http_response_code(401);
            print("DataNotFull");
        }
    } else {
        http_response_code(400);
        print("NoData");
    }
};
/* ******************************** */

function createJWT($userEmail, $signingKey, $expTime){
    $payload = array(
        "exp" => time() + $expTime,
        "iat" => time(),
        "Email" => $userEmail,
    );
    return JWT::encode($payload, $signingKey );
}
function reCreate($jwt){
    $tks = \explode('.', $jwt);
    if ( count($tks) != 3 ) { 
        print("invalid jwt");
        return null; 
    };
    $payload = json_decode( b64_decode( $tks[1] ));
    if (!isset($payload->Email) ){
        print("bad payload");
        return null;
    }
    return createJWT($payload->Email, $_ENV["JWT_SIGNING_KEY"], $_ENV["EXP_TIME"]);

}
function extractUserEmailFromJWT($jwt){
    $tks = \explode('.', $jwt);
    if ( count($tks) != 3 ) { 
        print("invalid jwt");
        return null; 
    };
    $payload = json_decode( base64_decode( $tks[1] ));
    if (!isset($payload->Email) ){
        print("bad payload");
        return null;
    }
    return $payload->Email;

}
function createRT($jwt, $signingKey){
    $data = explode(".", $jwt)[2];
    //$rt = "";
    //openssl_sign($data, $rt, $signingKey, "HS256");
    
    return base64_encode( hash_hmac("SHA256", $data, $signingKey) ) ;
}
function checkTokens($jwt, $rt){
    $jwtPart = explode(".", $rt)[2];
    return strcmp($jwtPart, $rt)? false :true ;
}
//example: pg_connect("host=postgres user=postgres password="password" port=5432 dbname=switter");
class DBHelper {
    private $dbConn=null;
    function __construct($connString){
        $this->$dbConn = pg_connect($connString);
    }
    function __destruct(){
        pg_close($this->$dbConn);
    }
    public function createUser($userName, $userEmail, $password, $userRT){
        $result = pg_query_params($this->$dbConn, 
            "INSERT INTO users(user_name,user_password,user_email,user_rt) VALUES($1, $2, $3, $4)",
            array($userName, $password, $userEmail, $userRT));
        return $result? true: false;
    }
    public function checkUserExist($userEmail){
        $result = pg_query_params($this->$dbConn, 
            "SELECT * FROM users WHERE user_email=$1 ;",
            array($userEmail));
        if (pg_fetch_all($result) != null) {
            return true;
        } else {
            return false;
        }
    }
    public function checkPassword($userEmail, $password){
        $query = pg_query_params($this->$dbConn, 
            "SELECT user_password FROM users WHERE user_email=$1 ;",
            array($userEmail));
        $result = pg_fetch_all($query);
        return password_verify($password, $result[0]["user_password"]);
    }
    public function checkRT($userEmail, $rt){
        $query = pg_query_params($this->$dbConn, 
            "SELECT user_rt FROM users WHERE user_email=$1 ;",
            array($userEmail)
        );
        $result = pg_fetch_all($query);
        if($result == null ) return false;
        return password_verify($rt, $result[0]["user_rt"]);
    }
    public function updateToken($userEmail, $rt){
        $result = pg_query_params($this->$dbConn, 
            "UPDATE USERS SET user_rt=$1 WHERE user_email=$2 ;",
            array( $rt, $userEmail));
            return $result? true: false;
    }
    public function deleteToken($userEmail){
        $result = pg_query_params($this->$dbConn, 
            "UPDATE USERS SET user_rt=0 WHERE user_email=$1 ;",
            array($userEmail));
            return $result? true: false;
    }
}
/* ************ */
$routes = [
    "/auth/login" => $LoginCallback,
    "/auth/register" => $RegisterCallback,
    "/auth/update" => $UpdateCallback,
    "/auth/delete" => $DeleteCallback,
];
//print_r( $_SERVER["REQUEST_URI"]);
if($_SERVER["REQUEST_URI"]){
    $routes[ $_SERVER["REQUEST_URI"] ]->__invoke("dude");
}
?>