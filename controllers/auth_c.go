package controllers
import (
	"github.com/gin-gonic/gin"	
	"strings"
	"encoding/base64"
	"net/http"
	"log"
	"time"
	"os"
	"strconv"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt"
	"encoding/json"
	"hris/models"
)
type TokenStruct struct {
	User models.User `json:"user"`
	Menu []models.MenuAccess1 `json:"menu"`
	Submenu []models.SubmenuAccess2 `json:"submenu"`
}
type AccessMenu struct {
	Menu []models.MenuAccess1 `json:"menu"`
	Submenu []models.SubmenuAccess2 `json:"submenu"`
}
type TokenStruct2 struct {
	User models.User `json:"user"`
}
func CheckEmailExist(c *gin.Context)([]models.User){
	var user models.User
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	return models.GetUser(user)
}
func CheckUsernameExist(c *gin.Context)([]models.User){
	var user models.User
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	return models.GetUser(user)
}
func CheckAPI(c *gin.Context)(a models.Application){
	tokenIn,_:=c.Cookie("token")
	
	var arr =c.Request.Header["Authorization"]
	if len(arr)>0{
		auths:=strings.Split(arr[0], " ")
		authType:=auths[0]
		authToken:=auths[1]
		if authType =="App"{
			app:=models.GetAPIKey(models.Application{Key:authToken,Typekey:authType})
			if app.Key!=""{
				currentTime := time.Now().Format("2006-01-02")
				dExp, _ := time.Parse("2006-01-02", app.Enddate)
				now,_:=time.Parse("2006-01-02", currentTime)
				if dExp.Unix()>=now.Unix(){
					a=app
					log.Println("API is valid. Welcome")
				}else{
					c.JSON(401, gin.H{"result":"API expired"})
				}
			}else{
				c.JSON(401, gin.H{"result":"Your API Key is not valid"})
			}
		}
	}else if tokenIn != ""{
		token, _:= jwt.Parse(tokenIn, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("TOKEN_SALT")), nil
			})
		if token.Valid{
			claim := claimJWT(token)
			a=models.GetAPIKey(models.Application{Id:claim.User.Idapplication})
		}else{
			// c.JSON(401, gin.H{"result":"Token is invalid"})
			c.Redirect(302,"/login/"+base64.StdEncoding.EncodeToString([]byte(c.Request.URL.Path)))
		}
	}else{
		// c.JSON(401, gin.H{"result":"Authorization is needed"})
		c.Redirect(302,"/login/"+base64.StdEncoding.EncodeToString([]byte(c.Request.URL.Path)))
		
	}
	a.Key="XXX"
	return a
}
func LoginWeb(c *gin.Context){
	username:=c.Request.FormValue("username")
	password:=c.Request.FormValue("password")
	rememberme:=c.Request.FormValue("rememberme")
	from:=c.Request.FormValue("from")

	expTime:=time.Now().Local().Add(time.Hour*8).Unix()
	if rememberme=="on"{
		expTime=time.Now().Local().Add(time.Hour*24).Unix()
	}
	
	if username != "" && password != ""{
		userData:=models.GetUserFull(models.User{Username:username})
		
		if userData.Username==""{
			c.Redirect(http.StatusFound,"/login/"+base64.StdEncoding.EncodeToString([]byte(from)))
		}else{
			
			if CheckPasswordHash(password, userData.Password)==true{
				userData.Password="***Mau ngintip password ya ???***"
				tokenStruct:=TokenStruct{User:userData}
				jsonData, _ := json.Marshal(tokenStruct)
				log.Println("jengggggggggggggggggggggggg",jsonData)

				generateJWT(jsonData,expTime,c)

				c.Redirect(http.StatusFound,from)
			}else{
				
				c.Redirect(http.StatusFound,"/login")
			}
		}
	}else{
		c.Redirect(http.StatusFound,"/login")
	}	
}
func LoginApp(c *gin.Context){
	var user models.User
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	var acc AccessMenu
	if user.Username != "" && user.Password != ""{
		userData:=models.GetUserFull(models.User{Username:user.Username})
		var menuData []models.MenuAccess1
		var menu2Data []models.SubmenuAccess2
		log.Println("userData",userData)
		if userData.Username==""{
			c.JSON(401, gin.H{"result":userData})
		}else if CheckPasswordHash(user.Password,userData.Password) == false{
			c.JSON(401, gin.H{"result":"Username or Password is wrong"})
		}else{
			menuData=models.GetMenuForLogin(userData.Idrole,userData.Idapplication)
			menu2Data=models.GetSubmenuForLogin(userData.Idrole,userData.Idapplication)
			acc= AccessMenu{Menu:menuData,Submenu:menu2Data}
			c.JSON(200, gin.H{"result":acc})
		}		
	}else{
		c.JSON(401, gin.H{"result":"Unauthorized"})
	}
}

func ClaimToken(c *gin.Context)(user models.User){
	tokenIn,_:=c.Cookie("token")
	token, err:= jwt.Parse(tokenIn, func(token *jwt.Token) (interface{}, error) {
				return []byte(os.Getenv("TOKEN_SALT")), nil
			})
	if token.Valid{
		claim := claimJWT(token)
		user=models.GetUserFull(models.User{Id:claim.User.Id})
		user.Password="***Mau ngintip password ya ???***"
		log.Println(user)
	}else{
		c.JSON(401, gin.H{"result":err})
	}
	return user
}
func EncodingBcrypt (p string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(p),1)
    return string(bytes), err
}
func CheckPasswordHash(password string, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    if err != nil {
    	log.Println(err)
    }
    return err == nil
}

func LogoutWeb(c *gin.Context){
	c.SetCookie("token", "xxx",-1, "/", os.Getenv("APP_HOST"), false, true)
	c.Redirect(http.StatusFound,"/")
}
func generateJWT(json []byte,expiredtime int64,c *gin.Context)(string){
	mySigningKey := []byte(os.Getenv("TOKEN_SALT"))
	type Claims struct {
		Json string `json:"data"`
		jwt.StandardClaims
	}
	claims := Claims{string(json),jwt.StandardClaims{ExpiresAt:expiredtime}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil{
		log.Println("generateJWT error : ",err)
	}
	c.SetCookie("token",ss, int(expiredtime), "/",os.Getenv("APP_HOST"), false, true)
	return ss
}
func claimJWT(token *jwt.Token)(TokenStruct){
	claims := token.Claims.(jwt.MapClaims)
	data:=claims["data"].(string)
	var t TokenStruct
	json.Unmarshal([]byte(data),&t)
	return t
}
// func CreateAPIKey(c *gin.Context,u models.User)(string){
// 	name:=c.Request.FormValue("name")
// 	datefrom:=c.Request.FormValue("datefrom")
// 	dateexp:=c.Request.FormValue("dateexp")
// 	if(name == "" || datefrom=="" || dateexp==""){
// 		return ""
// 	}
// 	dFrom, _ := time.Parse("2006-01-02", datefrom)
// 	dExp, _ := time.Parse("2006-01-02", dateexp)
// 	jsonUserData, _ := json.Marshal(u)
// 	hash,_:=EncodingBcrypt(string(jsonUserData))
// 	a:=models.ApiKey{Iduser:u.Id,Name:name,Type:"API Token",Validfrom:dFrom.Unix(),Validto:dExp.Unix(),Key:string(hash)}	
// 	res:=models.CreateApiKey(a)
// 	return string(res.Key)
// }
// func CreateBearerKey(c *gin.Context,u models.User)(string){
// 	name:=c.Request.FormValue("name")
// 	datefrom:=c.Request.FormValue("datefrom")
// 	dateexp:=c.Request.FormValue("dateexp")
// 	if(name == "" || datefrom=="" || dateexp==""){
// 		return ""
// 	}
// 	dFrom, _ := time.Parse("2006-01-02", datefrom)
// 	dExp, _ := time.Parse("2006-01-02", dateexp)

// 	mySigningKey := []byte(os.Getenv("TOKEN_SALT"))
// 	type Claims struct {
// 		Json string `json:"data"`
// 		jwt.StandardClaims
// 	}
// 	menuData:=models.GetMenuForLogin(u.Idrole)
// 	submenuData:=models.GetSubmenuForLogin(u.Idrole)

// 	data:=TokenStruct{User:u,Menu:menuData,Submenu:submenuData}
// 	xxx,_:=json.Marshal(data)
// 	claims := Claims{string(xxx),jwt.StandardClaims{ExpiresAt:dExp.Unix()}}
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	ss, err := token.SignedString(mySigningKey)
// 	if err != nil{
// 		log.Println("generateJWT error : ",err)
// 	}

// 	a:=models.ApiKey{Iduser:u.Id,Name:name,Type:"Bearer Token",Validfrom:dFrom.Unix(),Validto:dExp.Unix(),Key:string(ss)}	
// 	res:=models.CreateBearerKey(a)
// 	return string(res.Key)
// }
// func GetUserByJson(c *gin.Context)(models.User){
// 	var user models.User
// 	jsonData, _ := c.GetRawData()
// 	json.Unmarshal(jsonData,&user)
// 	u:=models.GetUser(user)
// 	u.Password="***Mau ngintip password ya ???***"
// 	return u
// }
func GetUserAccessMenu(c *gin.Context)(acc AccessMenu){
	var user models.User
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&user)
	if user.Idrole!=0{
		dataMenu1:=models.GetMenuForLogin(user.Idrole,user.Idapplication)
		dataMenu2:=models.GetSubmenuForLogin(user.Idrole,user.Idapplication)
		acc=AccessMenu{Menu:dataMenu1,Submenu:dataMenu2}
	}
	return acc
}
func GenerateAPIKey(c *gin.Context)(app models.Application){
	var a models.Application
	jsonData, _ := c.GetRawData()
	json.Unmarshal(jsonData,&a)
	if a.Id!=0{
		now:=strconv.FormatInt(time.Now().Unix(), 10)
		newkey,_:=EncodingBcrypt(now)
		app=models.UpdateApplication(models.Application{Id:a.Id,Key:newkey})
	}
	return app
}