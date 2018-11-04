package main

import (
	"fmt"
	"time"

	"github.com/OkumuraShintarou/peace/app"
	"github.com/OkumuraShintarou/peace/config"
	"github.com/OkumuraShintarou/peace/db"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.New()
	dbm := db.MustNewDB(cfg.DBUser, cfg.DBPass, cfg.DBName)
	app.Init(dbm, cfg)
}

func init() {
	s := `
	PPPPPPPPPPPPPPPPP   EEEEEEEEEEEEEEEEEEEEEE               AAA                  CCCCCCCCCCCCCEEEEEEEEEEEEEEEEEEEEEE
P::::::::::::::::P  E::::::::::::::::::::E              A:::A              CCC::::::::::::CE::::::::::::::::::::E
P::::::PPPPPP:::::P E::::::::::::::::::::E             A:::::A           CC:::::::::::::::CE::::::::::::::::::::E
PP:::::P     P:::::PEE::::::EEEEEEEEE::::E            A:::::::A         C:::::CCCCCCCC::::CEE::::::EEEEEEEEE::::E
  P::::P     P:::::P  E:::::E       EEEEEE           A:::::::::A       C:::::C       CCCCCC  E:::::E       EEEEEE
  P::::P     P:::::P  E:::::E                       A:::::A:::::A     C:::::C                E:::::E             
  P::::PPPPPP:::::P   E::::::EEEEEEEEEE            A:::::A A:::::A    C:::::C                E::::::EEEEEEEEEE   
  P:::::::::::::PP    E:::::::::::::::E           A:::::A   A:::::A   C:::::C                E:::::::::::::::E   
  P::::PPPPPPPPP      E:::::::::::::::E          A:::::A     A:::::A  C:::::C                E:::::::::::::::E   
  P::::P              E::::::EEEEEEEEEE         A:::::AAAAAAAAA:::::A C:::::C                E::::::EEEEEEEEEE   
  P::::P              E:::::E                  A:::::::::::::::::::::AC:::::C                E:::::E             
  P::::P              E:::::E       EEEEEE    A:::::AAAAAAAAAAAAA:::::AC:::::C       CCCCCC  E:::::E       EEEEEE
PP::::::PP          EE::::::EEEEEEEE:::::E   A:::::A             A:::::AC:::::CCCCCCCC::::CEE::::::EEEEEEEE:::::E
P::::::::P          E::::::::::::::::::::E  A:::::A               A:::::ACC:::::::::::::::CE::::::::::::::::::::E
P::::::::P          E::::::::::::::::::::E A:::::A                 A:::::A CCC::::::::::::CE::::::::::::::::::::E
PPPPPPPPPP          EEEEEEEEEEEEEEEEEEEEEEAAAAAAA                   AAAAAAA   CCCCCCCCCCCCCEEEEEEEEEEEEEEEEEEEEEE
	`
	fmt.Println(s)
	fmt.Println("Set Time Zone...", time.UTC)
	time.Local = time.UTC
}
