/*
使用者是否在遊戲中
他手上有哪兩張牌
*/
package DB
import (
	"os"
	"strconv"
	"github.com/line/line-bot-sdk-go/linebot"
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
/*
傳入使用者MID
回傳使用者是否正在遊戲
*/
var bot *linebot.Client
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
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
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
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var haveGame string
	var RID string
	var R string
	var GID string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT ID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&GID)
	db.QueryRow("SELECT RoomID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&haveGame)
	if haveGame == ""{
		db.Exec("INSERT INTO sql6131889.Game (GameName, RoomID, GameStatus, GameTokens, GamePlayer1, GameMaster, Cancel) VALUES (?, ?, ?, ?, ?, ?, ?)", "TexasPoker", RID, 100, 0, MID, "0", 0)
		
		db.Exec("INSERT INTO sql6131889.GameAction (MID, GameID, PlayerX, Action, Cancel) VALUE (?, ?, ?, ?, ?)", MID, GID, 20, , 0)
		bot.SendText([]string{MID}, "You created a new game")
	}else{
		bot.SendText([]string{MID}, "There is already a game in this room!!")
	}
	db.Close()
}
func InRoomJoinGame(MID string){
	strID := os.Getenv("ChannelID")
	numID, _ := strconv.ParseInt(strID, 10, 64) // string to integer
	bot, _ = linebot.NewClient(numID, os.Getenv("ChannelSecret"), os.Getenv("MID"))
	db,_ := sql.Open("mysql", os.Getenv("dbacc")+":"+os.Getenv("dbpass")+"@tcp("+os.Getenv("dbserver")+")/")
	var haveGame string
	var RID string
	var R string
	var GID string
	db.QueryRow("SELECT UserRoom FROM sql6131889.User WHERE MID = ?", MID).Scan(&R)
	db.QueryRow("SELECT ID FROM sql6131889.Room WHERE  RoomName = ?", R).Scan(&RID)
	db.QueryRow("SELECT ID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&GID)
	db.QueryRow("SELECT RoomID FROM sql6131889.Game WHERE RoomID = ?", RID).Scan(&haveGame)
	if haveGame == ""{
		bot.SendText([]string{MID}, "Please create a new game use instruction:\n!newgame")
	}else{
		var playerInGame string
		db.QueryRow("SELECT MID FROM sql6131889.GameAction WHERE MID = ?", MID).Scan(&playerInGame)
		if playerInGame == "" {
			row,_ := db.Query("SELECT PlayerX FROM sql6131889.GameAction WHERE GameID = ?", GID)
			var nextPlayer int
			for row.Next() {
				row.Scan(&nextPlayer)
			}
			nextPlayer = nextPlayer+1
		}else{
			nextPlayer = 50
		}
		if nextPlayer <= 29 {
			db.Exec("INSERT INTO sql6131889.GameAction (MID, GameID, PlayerX, Action, Cancel) VALUE (?, ?, ?, ?, ?)", MID, GID, 20, nextPlayer, 0)
			if nextPlayer == 21 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer2 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 22 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer3 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 23 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer4 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 24 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer5 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 25 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer6 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 26 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer7 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 27 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer8 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 28 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer9 = ? WHERE GID = ?", MID, GID)
			}else if nextPlayer == 29 {
				db.Exec("UPDATE sql6131889.Game SET GamePlayer10 = ? WHERE GID = ?", MID, GID)
			}
		}else if nextPlayer == 50 {
			bot.SendText([]string{MID}, "You are already in this game!!")
		}else{
			bot.SendText([]string{MID}, "Full of player in this room!!")
		}
	}
	db.Close()
}
