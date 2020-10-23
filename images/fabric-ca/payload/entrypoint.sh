#!/bin/bash

confcenter(){
    chmod +x /usr/local/bin/confcenterdownload
    confcenterdownload
    if [[ $? -ne 0 ]];then
        echo -e "Download configuration file error!\n"
        exit 11
    fi

}
confcenter
fabric-ca-server start -b admin:adminpw -p 7054 -d