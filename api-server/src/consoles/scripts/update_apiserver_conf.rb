require 'inifile'

conf = IniFile.load("../../conf/testing.conf")
conf['def_db']['db_port'] = 3306
conf['def_db']['db_host'] = 'rds2g2rdxxt711sworh8.mysql.rds.aliyuncs.com'
conf['def_db']['db_user'] = 'chinarun' 
conf['def_db']['db_pass'] = 'chEtHe7h' 
conf['def_db']['db_name'] = 'chinarun'
conf.write()
