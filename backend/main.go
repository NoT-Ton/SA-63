package main
 
import (
   "context"
   "log"
   "fmt"
 
   "github.com/NoT-Ton/app/controllers"
   _ "github.com/NoT-Ton/app/docs"
   "github.com/NoT-Ton/app/ent"
   "github.com/NoT-Ton/app/ent/role"
   "github.com/gin-contrib/cors"
   "github.com/gin-gonic/gin"
   _ "github.com/mattn/go-sqlite3"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"
)

type Users struct {
	User []User
}

type User struct {
    Email string
    Name  string
    Role int
}

type Roles struct {
	Role []Role
}

type Role struct {
	Name string
}

type Products struct {
	Product []Product
}

type Product struct {
	Name  string
	Cost int
}

type Customers struct {
	Customer []Customer
}

type Customer struct {
	Name  string
}

// @title SUT SA Example API
// @version 1.0
// @description This is a sample server for SUT SE 2563
// @termsOfService http://swagger.io/terms/
 
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
 
// @host localhost:8080
// @BasePath /api/v1
 
// @securityDefinitions.basic BasicAuth
 
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
 
// @securitydefinitions.oauth2.application OAuth2Application
// @tokenUrl https://example.com/oauth/token
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.implicit OAuth2Implicit
// @authorizationUrl https://example.com/oauth/authorize
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.password OAuth2Password
// @tokenUrl https://example.com/oauth/token
// @scope.read Grants read access
// @scope.write Grants write access
// @scope.admin Grants read and write access to administrative information
 
// @securitydefinitions.oauth2.accessCode OAuth2AccessCode
// @tokenUrl https://example.com/oauth/token
// @authorizationUrl https://example.com/oauth/authorize
// @scope.admin Grants read and write access to administrative information
func main() {
   router := gin.Default()
   router.Use(cors.Default())
 
   client, err := ent.Open("sqlite3", "file:ent.db?cache=shared&_fk=1")
   if err != nil {
       log.Fatalf("fail to open sqlite3: %v", err)
   }
   defer client.Close()
 
   if err := client.Schema.Create(context.Background()); err != nil {
       log.Fatalf("failed creating schema resources: %v", err)
   }
 
   v1 := router.Group("/api/v1")
   controllers.NewUserController(v1, client)
   controllers.NewRoleController(v1, client)
   controllers.NewProductController(v1, client)
   controllers.NewCustomerController(v1, client)
   controllers.NewBillController(v1, client)

  

    // Set Roles Data
    roles := Roles{
		Role: []Role{
			Role{"Librarian"},
            Role{"Seller"},
            Role{"Library Member"},
		},
	}
	for _, r := range roles.Role {
		client.Role.
			Create().
			SetRoleName(r.Name).
			Save(context.Background())
	}

	 // Set Users Data
	 users := Users{
		User: []User{
			User{"b6108526@g.sut.ac.th", "Warameth Nuipian",1},
			User{"me@example.com", "Name Surname",2},
			User{"Example@sa.com", "Ex System",2},
			User{"orange@sa.com", "ora System",3},
		},
	}

	for _, u := range users.User {

        r, err := client.Role.
			Query().
			Where(role.IDEQ(int(u.Role))).
            Only(context.Background())
        
        if err != nil {
			fmt.Println(err.Error())
			return
		}


		client.User.
			Create().
			SetUserEmail(u.Email).
            SetUserName(u.Name).
            SetRole(r).
			Save(context.Background())
	}

	// Set Product Data
	products := Products{
		Product: []Product{
			
			Product{"Paper", 1},
			Product{"Lab Cover", 10},
			Product{"Notebook", 30},
			Product{"MathBook", 100},
			Product{"ScienceBook", 120},
			Product{"EnglishBook", 165},	
		},
	}

	for _, p := range products.Product {
        client.Product.
            Create().
            SetProductName(p.Name).
            SetProductCost(p.Cost).
            Save(context.Background())
	}

	// Set Customers Data
	customers := Customers{
		Customer: []Customer{
			Customer{"-------------------"},
			Customer{"Alongkon Wanapradit"},
			Customer{"Nattanon Lakpakdee"},
		},
	}

	for _, c := range customers.Customer {

		client.Customer.
			Create().
			SetCustomerName(c.Name).
			Save(context.Background())
	}

 
   router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
   router.Run()
}
