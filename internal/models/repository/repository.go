package repository

// "github.com/nutikuli/internProject_backend/internal/services/file/entities"
// repositoryquery "github.com/nutikuli/internProject_backend/internal/services/file/repository/repository_query"

// func GetAdminuser(db *sql.DB) ([]*entities.Admin, error) {

// 	rows, err := db.Query(repositoryquery.SQL_get_account, "admin")
// 	if err != nil {
// 		log.Info(err)
// 		return nil, err
// 	}

// 	var admin []*entities.Admin
// 	err = scan.Row(&admin, rows)
// 	if err != nil {
// 		log.Errorf(" %#v", err)
// 		return nil, err
// 	}
// 	log.Infof("data %#v", admin)

// 	return admin, nil
// }

// func GetStoreuser(db *sql.DB) ([]*entities.Store, error) {

// 	rows, err := db.Query(repositoryquery.SQL_get_account, "store")
// 	if err != nil {
// 		log.Info(err)
// 		return nil, err
// 	}

// 	var store []*entities.Store
// 	err = scan.Row(&store, rows)
// 	if err != nil {
// 		log.Errorf(" %#v", err)
// 		return nil, err
// 	}
// 	log.Infof("data %#v", store)

// 	return store, nil
// }

// func GetCustomeruser(db *sql.DB) ([]*entities.Customer, error) {

// 	rows, err := db.Query(repositoryquery.SQL_get_account, "customer")
// 	if err != nil {
// 		log.Info(err)
// 		return nil, err
// 	}

// 	var customer []*entities.Customer
// 	err = scan.Row(&customer, rows)
// 	if err != nil {
// 		log.Errorf(" %#v", err)
// 		return nil, err
// 	}
// 	log.Infof("data %#v", customer)

// 	return customer, nil
// }
