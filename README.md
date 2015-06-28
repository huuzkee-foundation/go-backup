go-backup
=========


Input  options are : 

                      backup-path, (do we need multiple paths?)
                      
                      backup-frequency, 
                      
                      encryption,
                      
                      compression, 
                      
                      backup-server-location (ftp, s3, SoftLayer Object Storage),
                      
                      backup-runtine ??(do we need it)
                      
                      rollback-count(can we do manage that?)


The Main table/data structure/hash map will store these fields both at client side and serve side:

                      file full path,
                      
                      hash,
                      
                      hash_timestamp,  


Eveytime a new file is updated (if modification_timstamp>=hash_timestamp and hash_new!=has(in the table):-
 1. the file is the file is pushed into server
 2. for the first time the whole table/ds/hash from client will be replicated (bkp_data0)
 3. for the >first times every iteration of change .. the row corresponding updated  entries will be moved to backup_I where I is the iteration count. 
 4. server will store iteration count untill < rollback-count (if iteration count >rollback count we rewrite the main backup data table..update bkp_data0 to latest table on the client)
5.if user wants to rull back to backup_2..we superimpose bkp_0+bkup_1+bkup2(this will basically mean getting bckup_0 nad overwriting the latest updated/changed files until backup_2) and return the result to user.
