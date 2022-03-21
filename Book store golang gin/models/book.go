package models
import(
	"fmt"
	Config"book/config"

	_ "github.com/go-sql-driver/mysql"
)

func GetALLBook(book *[]Book)(err error){
	if err =Config.DB.Find(book).Error;err!=nil{
	return err

}
     return nil
}	 
func CreateABook(book *Book)(err error){
	if err =Config.DB.Create(book).Error;err !=nil{
		return err
	}
	return err
}
    func GetABook(book *Book, id string) (err error) {
	if err := Config.DB.Where("id = ?", id).First(book).Error; err != nil {
		return err
	}
	return nil
}

func UpdateABook(book *Book, id string) (err error) {
	fmt.Println(book)
	Config.DB.Save(book)
	return nil
}

func DeleteABook(book *Book, id string) (err error) {
	Config.DB.Where("id = ?", id).Delete(book)
	return nil
}

