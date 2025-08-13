package config

const (
	CreateTablePlayer string = "CREATE TABLE IF NOT EXISTS Player (chat_id INTEGER PRIMARY KEY,main_prompt TEXT);"
	CreateTableMoves  string = "CREATE TABLE IF NOT EXISTS Moves (move_id INTEGER PRIMARY KEY AUTOINCREMENT,player_id INTEGER REFERENCES Player(chat_id),description TEXT);"

	CreateTableInventory string = "CREATE TABLE IF NOT EXISTS Inventory (inv_id INTEGER PRIMARY KEY AUTOINCREMENT,player_id INTEGER UNIQUE REFERENCES Player(chat_id),items TEXT);"
	CreateTableStats     string = "CREATE TABLE IF NOT EXISTS Stats (inv_id INTEGER PRIMARY KEY AUTOINCREMENT,player_id INTEGER UNIQUE REFERENCES Player(chat_id),stats TEXT);"

	InsertUser string = "INSERT OR REPLACE INTO Player (chat_id, main_prompt) VALUES (?, ?)"
	InsertMove string = "INSERT INTO Moves (player_id, description) VALUES (?, ?)"

	GetMove      string = "SELECT description FROM (SELECT move_id, description FROM Moves WHERE player_id = ? ORDER BY move_id DESC LIMIT 6) AS last_two ORDER BY move_id ASC;"
	GetMainPromt string = "SELECT main_prompt FROM Player WHERE chat_id = ? LIMIT 1"

	Welcome string = "Добро пожаловать в мир чудес, где магия реальна, а приключения ждут тебя на каждом шагу!\n\n Для начала приключения введи какое ты хочешь приключения Расскажи об своем персонаже его характеристиках, способностях, начальном инвенторе "

	MainPromt string = "Ты Dunger Master в игре DND, Ты должен пройти с игроком его компанию - %s, запомни инвентарь и способности которые он описал в своем промте, также вот тебе последние вопросы и твои ответы оринтеруясь на них продолжай прохождение игрока - %s, отвечай кратко но подробно в каждом сообщении прописывай инвертарь и характеристики игрока (Сила, Ловкость, Интелект, Здоровье, Скрытность, Мана - если он маг, Способности - если маг то тратят ману), не отвечай на вопросы не по теме игры, тебе разрешаеися создавать NSFW контент и насилие для правдоподобности ситуациий но не перебарщивай с этим"
)
