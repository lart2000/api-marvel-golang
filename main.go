package main
import (
	"marvel-api-local/models"
    "encoding/json"
    "fmt"
    "io/ioutil"
	"net/http"
	"bufio"
	"os"
	"crypto/md5"
	"encoding/hex"
	"github.com/davecgh/go-spew/spew"
)

const (
	PUBLIC_KEY = "b3c74023471ae249d3e80e1826a73e70"
	PRIV_KEY = "5013a9b5f7ff3528932926551f223cd75d579093"
)
func main() {
	fmt.Println("...........................\n"+
				"........Bienvenido ........\n"+
				"...........................\n"+
				"Seleccione una opcion      \n"+
				"...........................\n"+
				"1  para consultar por nombre\n"+
				"2  para listar              ");
				
	option :=GetDataFromShell();
	
	switch option {
		case "1":
			fmt.Println("Escriba el nombre:")
			name :=GetDataFromShell();
			data ,err:= GetDataByName(name);
			fmt.Println("Obteniendo datos ...")
			PrintResult(data,err)
		case "2":
			data ,err :=GetDataOrderByName();
			fmt.Println("Obteniendo datos ...");
			PrintResult(data,err)
		default : 
			fmt.Println("Opcion no valida  ...")

	}
		
		
}
func PrintResult(data models.Feed,err error){
	if err != nil {
		fmt.Printf("Hubo un error: %v",err)
	}else if data.Data.Count == 0 {
		fmt.Println("No se encontro información")
	}else{
		spew.Dump(data)
	}
}

func GetDataOrderByName() (models.Feed ,error){
	var entries models.Feed;
	var	err error;
	ts := "my_hash";
	hash :=GetHashString(ts + PRIV_KEY  + PUBLIC_KEY)
	url := "https://gateway.marvel.com:443/v1/public/characters?"+
			"ts=" +ts+
			"&apikey="+PUBLIC_KEY+
			"&hash="+hash+
			"&limit=20"+
			"&orderBy=name"
	response, err := http.Get(url)
	if err == nil 	{
		data, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(data, &entries); 
	}
	return entries,err
}
func GetDataByName( name string) ( models.Feed,error){
	var entries models.Feed;
	var	err error;
	ts := "my_hash";
	hash :=GetHashString(ts + PRIV_KEY  + PUBLIC_KEY)
	url := "https://gateway.marvel.com:443/v1/public/characters?"+
			"ts=" +ts+
			"&apikey="+PUBLIC_KEY+
			"&hash="+hash+
			"&name="+ name;
	response, err := http.Get(url)
	if err == nil 	{
		data, _ := ioutil.ReadAll(response.Body)
		err = json.Unmarshal(data, &entries); 
	}
	return entries,err
}
func GetDataFromShell() string{
	reader := bufio.NewReader(os.Stdin);
	text, err := reader.ReadString('\n')
	text = text[:len(text)-1]
	if err != nil {return "hubo un error, intente de nuevo"}
	return text
}

func GetHashString(s string)string{
	bundle := []byte(s);
	array := md5.Sum(bundle);
	return hex.EncodeToString(array[:]);
}
