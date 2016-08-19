/*
使用者是否在遊戲中
他手上有哪兩張牌
*/
package DB
import(
	"os"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
/*
傳入使用者MID
回傳使用者是否正在遊戲
*/
func UserGamming(MID string) bool{
	var GameID int
	GameID = 0;
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	db.QueryRow("SELECT GameID FROM sql6131889.GameAction WHERE MID = ? and Cancel = 0", MID ).Scan(&GameID)
	if GameID == 0{
		return false
	}else{
		return true
	}


}
func InRoomInst(MID string){
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var haveGame string
	var RID string
	var R string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.User WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT RoomID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&haveGame)
	if haveGame == ""{
		bot.SendText([]string{MID}, "You can use these instruction:\n!leavechatroom\n!newgame")
	}else{
		bot.SendText([]string{MID}, "You can use these instruction:\n!leavechatroom")
	}
	db.Close()
}
func InRoomNewGame(MID string){
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var haveGame string
	var RID string
	var R string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT RoomID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&haveGame)
	if haveGame == ""{
		db.Exec("INSERT INTO sql6131889.Game (GameName, RoomID, GameStatus, GameTokens, GamePlayer1, GameMaster, Cancel) VALUES (?, ?, ?, ?, ?, ?, ?)", "TexasPoker", RID, 100, 0, info[0].MID, "0", 0)
		bot.SendText([]string{content.From}, "You created a new game")
	}else{
		bot.SendText([]string{content.From}, "There is already a game in this room!!")
	}
	db.Close()
}
func InRoomJoinGame(MID string){
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var haveGame string
	var RID string
	var R string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT RoomID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&haveGame)
	if haveGame == ""{
		bot.SendText([]string{content.From}, "Please create a new game")
	}else{
		db.Exec("UPDATE sql6131889.Game SET UserStatus = ? WHERE MID = ?", 10, MID)
	}
	db.Close()
}
