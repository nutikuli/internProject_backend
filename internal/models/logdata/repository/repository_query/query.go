package repository_query

// Create log

var SQL_insert_logdata = `INSERT INTO Log (fullname,menuRequest,actionRequest,createAt) VALUE(? , ? , ? , ?) ;` 


//Get log 

var SQL_get_logdata = `SELECT * FROM Log ;` 
