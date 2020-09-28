<?php
require './vendor/autoload.php';
use \Firebase\JWT\JWT;


/* ************ */
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
            if ( $dbHelper->userCheck($rbody["userEmail"]) ){
                if ($dbHelper->userPasswordCheck($rbody["userEmail"] , $rbody["password"] )) {
                    $jwt = createJWT($rbody["userEmail"], $_ENV["JWT_SIGNING_KEY"], $_ENV["EXP_TIME"]);
                    $rt = createRT($jwt, $_ENV["RT_SIGNING_KEY"]);
                    $encodedRT = password_hash($rt, PASSWORD_DEFAULT);
                    $dbHelper->userTokensUpdate($rbody["userEmail"], $encodedRT );
                    $responce = array(
                        "jwt" => $jwt,
                        "rt" => $rt,
                    );
                    print(json_encode($responce));
                }   else {
                    http_response_code(401);
                    print("password not match");    
                }
            } else {
                http_response_code(401);
                print("user not exist");
            }
        } else {
            http_response_code(400);
            print("auth data not full");
        }
    } else {
        http_response_code(400);
        print("auth data not full");
    }
    //print( $_ENV["SIGNING_KEY"]) ;
};
$RegisterCallback = function($val) {
    header('Content-Type: application/json');
    //echo "CreateCallback! ".$val."\n";
    $dbHelper = new DBHelper("host=postgres port=5432 user=postgres password=password port=5432 dbname=switter");
    $rbody = json_decode( file_get_contents('php://input'),true ) ;
    $errno = json_last_error();
    if ( $errno == null ){
        if( $rbody["userEmail"] != null && $rbody["password"] != null && $rbody["userName"] != null ) {
            // user existing db check
            
            if (!$dbHelper->userCheck($rbody["userEmail"]  )) {
                $jwt = createJWT($rbody["userEmail"], $_ENV["JWT_SIGNING_KEY"], $_ENV["EXP_TIME"]);
                $rt = createRT($jwt, $_ENV["RT_SIGNING_KEY"]);
                $encodedRT = password_hash($rt, PASSWORD_DEFAULT);
                $encodedPassword = password_hash($rbody["password"], PASSWORD_DEFAULT);
                $dbHelper->userRegister($rbody["userName"],$rbody["userEmail"],$encodedPassword, $encodedRT );
                $responce = array(
                    "jwt" => $jwt,
                    "rt" => $rt,
                );
                print(json_encode($responce));
            }   else {
                http_response_code(401);
                print("user exist");    
            }
        } else {
            http_response_code(400);
            print("auth data not full");
        }
    } else {
        http_response_code(400);
        print("auth data not full");
    }
};
$EditCallback = function($val) {
    //echo "EditCallback ".$val."!\n";
    print("nothing");
};
$UpdateCallback = function($val) {
    //echo "EditCallback ".$val."!\n";
    print("nothing");
};
$DeleteCallback = function($val) {
    //echo "DeleteCallback ".$val."!\n";
    print_r($_POST);
};
/**/ 

function createJWT($userEmail, $signingKey, $expTime){
    $payload = array(
        "StandardClaims" => array(
            "exp" => time() * $expTime,
            "iat" => time(),
        ),
        "userEmail" => $userEmail,
    );
    return JWT::encode($payload, $signingKey );
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
    public function userRegister($userName, $userEmail, $password, $userRT){
        $result = pg_query_params($this->$dbConn, 
            "INSERT INTO users(user_name,user_password,user_email,user_rt) VALUES($1, $2, $3, $4)",
            array($userName, $password, $userEmail, $userRT));
        return $result? true: false;
    }
    public function userCheck($userEmail){
        $result = pg_query_params($this->$dbConn, 
            "SELECT * FROM users WHERE user_email=$1 ;",
            array($userEmail));
        if (pg_fetch_all($result) != null) {
            return true;
        } else {
            return false;
        }
    }
    public function userPasswordCheck($userEmail, $password){
        $query = pg_query_params($this->$dbConn, 
            "SELECT user_password FROM users WHERE user_email=$1 ;",
            array($userEmail));
        $result = pg_fetch_all($query);
        return password_verify($password, $result[0]["user_password"]);
    }
    public function userTokensUpdate($userEmail, $rt){
        $result = pg_query_params($this->$dbConn, 
            "UPDATE USERS SET user_rt=$1 WHERE user_email=$2 ;",
            array( $rt, $userEmail));
            return $result? true: false;
    }
    public function userTokensDelete($userEmail){
        $result = pg_query_params($this->$dbConn, 
            "UPDATE USERS SET user_rt=0 WHERE user_email=$2 ;",
            array( $rt, $userEmail));
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