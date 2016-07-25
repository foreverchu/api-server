#! /bin/bash

mysqldump -uchinarun -p chinarun --no-data | sed 's/ AUTO_INCREMENT=[0-9]*//g' > chinarun_api.sql
