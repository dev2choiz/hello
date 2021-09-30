package diff

import (
	"fmt"
	"github.com/dev2choiz/hello/pkg/models"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
)

const migDir = "cmd/functions/migration/files"

// Generate generate a sql file migration
func Generate(db *gorm.DB) error {
	migLogger := newMigLogger()

	err := db.
		Session(&gorm.Session{Logger: migLogger, DryRun: true}).
		AutoMigrate(&models.Data{})
	if err != nil {
		return err
	}

	f, err := createFileMigration(migLogger.Sql)
	if err != nil {
		panic(err)
	}

	log.Println(f, "created")

	return nil
}

func createFileMigration(sqls []string) (string, error) {

	ver := getNextVersion(migDir)
	start := fmt.Sprintf(`package %s

import (
	"fmt"
	"github.com/go-pg/migrations/v8"
)

func init() {
	migrations.MustRegisterTx(func(db migrations.DB) error {
		var err error
`, path.Base(migDir))

	end := `
	}, func(db migrations.DB) error {
		return nil
	})
}`

	filename := filepath.Join(migDir, fmt.Sprintf("%d_new_migration.go", ver))
	//filename = filepath.Join(migDir, "1_init.go") // tmp @todo remove
	file, err := os.Create(filename)
	if err != nil {
		file.Close()
		panic(err)
	}
	_ = file.Truncate(0)

	_, err = file.Write([]byte(start))
	if err != nil {
		return "", err
	}

	// write sql
	for i, sql := range sqls {

		_, err = file.Write([]byte(`
		fmt.Println(` +  "`execute : " +  sql +  "`" +  `)
		_, err = db.Exec(` + "`" + sql + "`" + `)`))

		if i < len(sqls) - 1 {
			_, err = file.Write([]byte(`
		if err != nil {
			return err
		}
`))
			if err != nil {
				return "", err
			}
		} else {
			_, err = file.Write([]byte("\n\t\treturn err\n"))
			if err != nil {
				return "", err
			}
		}
	}

	_, err = file.Write([]byte(end))
	if err != nil {
		return "", err
	}

	return filename, nil
}

// getNextVersion return the next migration number
func getNextVersion(dir string) int {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}

	max := 0
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		str := strings.Split(f.Name(), "_")[0]
		num, err := strconv.Atoi(str)
		if err != nil {
			continue
		}
		if num > max {
			max = num
		}
	}
	return max + 1
}
