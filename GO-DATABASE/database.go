package main

// CREATE
func createProduct(product *Product) error {
	_,err := db.Exec(
		"INSERT INTO public.products(name, price) VALUES ($1, $2)",
		product.Name,
		product.Price,
	)
	return err
}

//GET
func getProduct(id int ) (Product,  error ) {
	var p Product 
	row := db.QueryRow(
		"SELECT id,name,price FROM products WHERE id=$1;",id,
	)
	err := row.Scan(&p.ID , &p.Name , &p.Price)

	if err != nil {
		return Product{} , err
	}

	return p , nil
}

//GET ALL 
func getProducts() ([]Product ,  error) {
	rows , err := db.Query("SELECT id , name ,price FROM products")
	if err != nil {
		return nil, err
	}
	var products []Product

	for rows.Next() {
		var p Product
		err := rows.Scan(&p.ID , &p.Name , &p.Price)

		if err != nil {
			return nil , err
		}
		products = append(products, p)
	}
	if err = rows.Err(); err != nil {
		return nil ,err
	}
	return products , nil
}

//UPDATE
func updateproduct(id int  , product *Product) (Product , error) {

	var p Product
	row := db.QueryRow(
		"UPDATE public.products SET  name=$1, price= $2  WHERE id=$3 RETURNING id , name ,price ;", product.Name , product.Price , id,
	)

	err := row.Scan(&p.ID , &p.Name , &p.Price)

	if err != nil {
		return Product{} , err
	}

	return p ,err
}

//DELETE
func deleteProduct(id int ) error {
	_, err := db.Exec(
		"DELETE FROM products WHERE id=$1;",id,
	) 
	return err
}


	// {/* GET */}

	// product , err := getProduct(4)
	// fmt.Println("Get Success" , product)

	//GET ALL PRODUCT 
	// products , err := getProducts()
	// fmt.Println(products)

	// {/* CREATE */}

	//  err = createProduct(&Product{Name: "BBB" , Price: 123})
	//  if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Create Success" )

	// {/* UPDATE */}

	// product,err = updateproduct(1,&Product{Name: "bbbbb" , Price: 212})
	// fmt.Println("Update Success" , product)

	// {/* DELETE */}

	// err = deleteProduct(0)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Delete Successful !" )

