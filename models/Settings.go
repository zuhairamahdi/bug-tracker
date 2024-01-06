package models

type Settings struct {
	Id int `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Settings json.RawMessage `json:"settings"`
	Name string `json:"name"`

}

//Default settings for Roles and Boards Role
var DefaultSettings = []Settings{
	{
		Settings: json.RawMessage(`{
			"roles": [
				{
					"name": "Admin",
					"permissions": [
						"createBoard",
						"deleteBoard",
						"updateBoard",
						"createRole",
						"deleteRole",
						"updateRole",
						"createTask",
						"deleteTask",
						"updateTask",
						"createUser",
						"deleteUser",
						"updateUser",
						"createComment",
						"deleteComment",
						"updateComment",
						"createColumn",
						"deleteColumn",
						"updateColumn",
						"createBoardRole",
						"deleteBoardRole",
						"updateBoardRole"
					]
				},
				{
					"name": "User",
					"permissions": [
						"createTask",
						"deleteTask",
						"updateTask",
						"createComment",
						"deleteComment",
						"updateComment"
					]
				},
				{
					"name": "Project Manager",
					"permissions": [
						"createTask",
						"deleteTask",
						"updateTask",
						"createComment",
						"deleteComment",
						"updateComment",
						"createColumn",
						"deleteColumn",
						"updateColumn"
					]
				},
				{
					"name": "Developer",
					"permissions": [
						"createTask",
						"deleteTask",
						"updateTask",
						"createComment",
						"deleteComment",
						"updateComment"
					]
				},
				{
					"name": "QA",
					"permissions": [
						"createTask",
						"deleteTask",
						"updateTask",
						"createComment",
						"deleteComment",
						"updateComment"
					]
				},
				{
					"name": "Board Maker",
					"permissions": [
						"createTask",
						"deleteTask",
						"updateTask",
						"createComment",
						"deleteComment",
						"updateComment",
						"createColumn",
						"deleteColumn",
						"updateColumn",
						"createBoardRole",
						"deleteBoardRole",
						"updateBoardRole"
						"createBoard",
						"deleteBoard",
						"updateBoard",
						]
				}
			],
			"boardRoles": [
				{
					"name": "Admin",
					"description": "Admin of the board",
					"permissions": [
						"createTask",
						"deleteTask",
						"updateTask",
						"createComment",
						"deleteComment",
						"updateComment",
						"createColumn",
						"deleteColumn",
						"updateColumn",
						"createBoardRole",
						"deleteBoardRole",
						"updateBoardRole"
					]
				},
				{
					"name": "User",
					"description": "User of the board",
					"permissions": [
						"createTask",
						"deleteTask",
						"updateTask",
						"createComment",
						"deleteComment",
						"updateComment"
					]
				}
			]
		}`),
		Name: "Default",
	},
	{
		Settings: json.RawMessage(`{
			"roles": [
				{
					"name": "Admin",
					"permissions": [
						"createBoard",
						"deleteBoard",
						"updateBoard",
						"createRole",
						"deleteRole",
						]
				},
				]
			}`),
	}