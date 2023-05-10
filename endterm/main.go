package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
)

type server struct {
	db *sql.DB
}

type product struct {
	Id    int
	Name  string
	Price int
}

type Order struct {
	Phone   string
	Product int
}

const (
	host     = "localhost"
	port     = 5433
	user     = "postgres"
	password = "12345"
	dbname   = "wedding"
)

func connect() *server {
	dbconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dbconn)
	if err != nil {
		log.Fatal(err)
	}
	return &server{db: db}
}

func (s *server) services(w http.ResponseWriter, r *http.Request) {
	result, err := s.db.Query("select * from products;")
	if err != nil {
		log.Fatal("query", err)
	}
	var products []product
	for result.Next() {
		var p product
		result.Scan(&p.Id, &p.Name, &p.Price)
		products = append(products, p)
	}
	t, err := template.ParseFiles("static/html/services.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, products)

}

func (s *server) reg(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		phone := r.FormValue("phone")
		address := r.FormValue("address")
		email := r.FormValue("email")
		pass := r.FormValue("psw")
		if _, err := s.db.Exec("insert into users(firstname, lastname, phone, address, email, password) values($1, $2, $3, $4, $5, $6)", firstname, lastname, phone, address, email, pass); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/auth", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("static/html/register.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, nil)
}

func (s *server) auth(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		phone := r.FormValue("phone")
		pass := r.FormValue("psw")
		var passCheck string
		if err := s.db.QueryRow("select password from users where phone=$1", phone).Scan(&passCheck); err != nil {
			log.Fatal(err)
		}
		if pass != passCheck {
			fmt.Print("Incorrect password!")
			http.Redirect(w, r, "/auth", http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("static/html/auth.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, nil)
}

func admin(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("static/html/admin.html")
	t.Execute(w, nil)
}

func (s *server) createWedding(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		price := r.FormValue("price")
		if _, err := s.db.Exec("insert into products(name, price) values($1, $2)", name, price); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/services", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("static/html/create.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, nil)
}

func (s *server) updateWedding(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		name := r.FormValue("name")
		price := r.FormValue("price")
		if _, err := s.db.Exec("update products set name=$1, price=$2 where id=$3", name, price, id); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/services", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("static/html/update.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, nil)
}

func (s *server) deleteWedding(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		name := r.FormValue("name")
		if _, err := s.db.Exec("delete from products where name=$1", name); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/services", http.StatusSeeOther)
		return
	}
	t, _ := template.ParseFiles("static/html/delete.html")
	t.Execute(w, nil)
}

func (s *server) updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		firstname := r.FormValue("firstname")
		lastname := r.FormValue("lastname")
		phone := r.FormValue("phone")
		address := r.FormValue("address")
		email := r.FormValue("email")
		pass := r.FormValue("psw")
		if _, err := s.db.Exec("update users set firstname=$1, lastname=$2, phone=$3, address=$4, email=$5, password=$6 where id=$7", firstname, lastname, phone, address, email, pass, id); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("static/html/updateUser.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, nil)
}

func (s *server) deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		if _, err := s.db.Exec("delete from users where id=$1", id); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/admin", http.StatusSeeOther)
		return
	}
	t, _ := template.ParseFiles("static/html/deleteUser.html")
	t.Execute(w, nil)
}

func order(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	fmt.Print(id)
	data := map[string]interface{}{"id": id}
	t, _ := template.ParseFiles("static/html/order.html")
	t.Execute(w, data)
}

func (s *server) orderFinal(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	phone := r.FormValue("phone")
	if _, err := s.db.Exec("insert into orders(phone, productid) values($1, $2)", phone, id); err != nil {
		log.Fatal(err)
	}
	http.Redirect(w, r, "/services", http.StatusSeeOther)
	return
}

func (s *server) orders(w http.ResponseWriter, r *http.Request) {
	result, err := s.db.Query("select * from orders;")
	if err != nil {
		log.Fatal("query", err)
	}
	var orders []Order
	for result.Next() {
		var o Order
		result.Scan(&o.Phone, &o.Product)
		orders = append(orders, o)
	}
	t, err := template.ParseFiles("static/html/orders.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, orders)

}

func (s *server) updateOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		phone := r.FormValue("phone")
		nid := r.FormValue("nid")
		nphone := r.FormValue("nphone")
		if _, err := s.db.Exec("update orders set productid=$1, phone=$2 where productid=$3 and phone=$4", nid, nphone, id, phone); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/ao", http.StatusSeeOther)
		return
	}
	t, err := template.ParseFiles("static/html/updateOrder.html")
	if err != nil {
		log.Fatal("parse", err)
	}
	t.Execute(w, nil)
}

func (s *server) deleteOrder(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		phone := r.FormValue("phone")
		if _, err := s.db.Exec("delete from orders where productid=$1 and phone=$2", id, phone); err != nil {
			log.Fatal(err)
		}
		http.Redirect(w, r, "/ao", http.StatusSeeOther)
		return
	}
	t, _ := template.ParseFiles("static/html/deleteOrder.html")
	t.Execute(w, nil)
}

func main() {
	s := connect()
	defer s.db.Close()
	fileServer := http.FileServer(http.Dir("./static/"))
	http.Handle("/", fileServer)
	http.HandleFunc("/services", s.services)
	http.HandleFunc("/reg", s.reg)
	http.HandleFunc("/auth", s.auth)
	http.HandleFunc("/admin", admin)
	http.HandleFunc("/cw", s.createWedding)
	http.HandleFunc("/uw", s.updateWedding)
	http.HandleFunc("/dw", s.deleteWedding)
	http.HandleFunc("/du", s.deleteUser)
	http.HandleFunc("/uu", s.updateUser)
	http.HandleFunc("/ao", s.orders)
	http.HandleFunc("/uo", s.updateOrder)
	http.HandleFunc("/do", s.deleteOrder)
	http.HandleFunc("/buy", order)
	http.HandleFunc("/orderfinal", s.orderFinal)
	http.ListenAndServe(":8000", nil)
}
