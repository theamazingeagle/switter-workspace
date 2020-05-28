#!/bin/bash
npm cache verify
npm install
npm install -g @vue/cli @vue/cli-service
npm install --save vue-axios axios
npm run serve
