package repository_query

// get bank
var SQL_get_banks = `SELECT * FROM Bank;`

// get bank by store id
var SQL_get_banks_by_store_id = `SELECT * FROM Bank WHERE storeId = ?;`

// get bank by id
var SQL_get_bank_by_id = `SELECT * FROM Bank WHERE id = ?;`

// insert bank
var SQL_insert_bank = `INSERT INTO Bank (name,accNumber,accName,avatarUrl,status,storeId) VALUE(?,?,?,?,?,?);`

var SQL_delete_bank_by_id = `DELETE FROM Bank WHERE id = ?;`

var SQL_update_bank_by_id = `UPDATE Bank SET name = ?, accNumber = ?, accName = ?, avatarUrl = ?, status = ?, storeId = ? WHERE id = ?;`

var SQL_get_bank_by_order_id = `SELECT b.* FROM Bank b JOIN Order o ON b.id = o.bankId WHERE o.id = ?;`
