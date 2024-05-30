package repository_query

// get bank
var SQL_get_bank = `SELECT id,name,accNumber,accName,status,avatarUrl,storeId FROM Bank;`

// get bank by store id
var SQL_get_bank_by_store_id = `SELECT id,name,accNumber,accName,status,avatarUrl,storeId FROM Bank WHERE storeId = ?;`

// get bank by id
var SQL_get_bank_by_id = `SELECT id,name,accNumber,accName,status,avatarUrl,storeId FROM Bank WHERE id = ?;`

//insert bank
var SQL_insert_bank = `INSERT INTO Bank (name,accNumber,accName,status,avatarUrl,storeId) VALUE(?,?,?,?,?,?);`
