## Golangでの画面作成、編集ウェブアプリのアイディア

### 要件

- 画面ごとに項目定義を登録。
- ホーム画面では、新規登録ボタンと登録済みの一覧をテーブル形式で表示。
- 編集画面では、項目定義の追加と登録済み項目定義の一覧を表示し、編集と削除が可能。
- インターフェース: Bootstrap
- 作成、編集処理: Golangサーバー側
- データの永続化にSQLiteを使用。

### 構成

1. **ホーム画面 (Bootstrap)**
   - 新規登録ボタン
   - 登録済みの一覧をテーブル形式で表示（列：画面名）
   - 画面名はリンクとして表示し、クリックすると編集画面へ遷移

2. **編集画面 (Bootstrap)**
   - 項目定義追加フォーム
   - 登録済み項目定義の一覧をテーブル形式で表示（列：項目名、アクション）
   - アクションセルには編集ボタンと削除ボタンを表示
   - 編集ボタンをクリックすると、項目定義のタイプごとに編集項目を動的に表示
   - 削除ボタンをクリックすると、削除確認メッセージを表示し、確認OKで削除

### ホーム画面の例 (HTML)

```html
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>ホーム画面</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
<div class="container">
    <h1>ホーム画面</h1>
    <button class="btn btn-primary" onclick="location.href='/components/new'">新規登録</button>
    <table class="table mt-4">
        <thead>
            <tr>
                <th>画面名</th>
            </tr>
        </thead>
        <tbody id="componentsList">
            <!-- 動的に追加される -->
        </tbody>
    </table>
</div>
<script>
    // 画面ロード時に登録済みの一覧を取得して表示
    fetch('/components')
        .then(response => response.json())
        .then(data => {
            const componentsList = document.getElementById('componentsList');
            data.forEach(component => {
                const tr = document.createElement('tr');
                const td = document.createElement('td');
                const a = document.createElement('a');
                a.href = `/components/${component.id}`;
                a.textContent = component.name;
                td.appendChild(a);
                tr.appendChild(td);
                componentsList.appendChild(tr);
            });
        });
</script>
</body>
</html>
```

### 編集画面の例 (HTML)

```html
<!DOCTYPE html>
<html lang="ja">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
    <title>編集画面</title>
    <link href="https://stackpath.bootstrapcdn.com/bootstrap/4.3.1/css/bootstrap.min.css" rel="stylesheet">
</head>
<body>
<div class="container">
    <h1>編集画面</h1>
    <form id="componentForm">
        <div class="form-group">
            <label for="id">ID</label>
            <input type="text" class="form-control" id="id" required>
        </div>
        <div class="form-group">
            <label for="label">ラベル</label>
            <input type="text" class="form-control" id="label" required>
        </div>
        <div class="form-group">
            <label for="type">タイプ</label>
            <select class="form-control" id="type" required>
                <option value="">選択してください</option>
                <option value="text">テキスト</option>
                <option value="textarea">テキストエリア</option>
                <option value="dropdown">ドロップダウンリスト</option>
                <option value="radio">ラジオボタンリスト</option>
                <option value="button">ボタン</option>
            </select>
        </div>
        <div class="form-group">
            <label for="maxLength">最大長</label>
            <input type="number" class="form-control" id="maxLength">
        </div>
        <div class="form-group">
            <label for="required">必須</label>
            <select class="form-control" id="required" required>
                <option value="true">はい</option>
                <option value="false">いいえ</option>
            </select>
        </div>
        <div class="form-group">
            <label for="options">オプション</label>
            <input type="text" class="form-control" id="options">
        </div>
        <div class="form-group">
            <label for="handler">イベントハンドラー</label>
            <input type="text" class="form-control" id="handler">
        </div>
        <button type="submit" class="btn btn-primary">保存</button>
    </form>
    <h2 class="mt-4">登録済み項目定義</h2>
    <table class="table mt-2">
        <thead>
            <tr>
                <th>項目名</th>
                <th>アクション</th>
            </tr>
        </thead>
        <tbody id="itemList">
            <!-- 動的に追加される -->
        </tbody>
    </table>
</div>
<script>
    // 編集画面のロード時に項目定義を取得して表示
    const urlParams = new URLSearchParams(window.location.search);
    const id = urlParams.get('id');

    fetch(`/components/${id}`)
        .then(response => response.json())
        .then(data => {
            document.getElementById('id').value = data.id;
            document.getElementById('label').value = data.label;
            document.getElementById('type').value = data.type;
            document.getElementById('maxLength').value = data.maxLength;
            document.getElementById('required').value = data.required;
            document.getElementById('options').value = data.options;
            document.getElementById('handler').value = data.handler;
        });

    fetch(`/components/${id}/items`)
        .then(response => response.json())
        .then(data => {
            const itemList = document.getElementById('itemList');
            data.forEach(item => {
                const tr = document.createElement('tr');
                const tdName = document.createElement('td');
                tdName.textContent = item.label;
                const tdActions = document.createElement('td');
                const editButton = document.createElement('button');
                editButton.textContent = '編集';
                editButton.classList.add('btn', 'btn-secondary', 'mr-2');
                editButton.onclick = () => editItem(item.id);
                const deleteButton = document.createElement('button');
                deleteButton.textContent = '削除';
                deleteButton.classList.add('btn', 'btn-danger');
                deleteButton.onclick = () => deleteItem(item.id);
                tdActions.appendChild(editButton);
                tdActions.appendChild(deleteButton);
                tr.appendChild(tdName);
                tr.appendChild(tdActions);
                itemList.appendChild(tr);
            });
        });

    function editItem(itemId) {
        fetch(`/components/${id}/items/${itemId}`)
            .then(response => response.json())
            .then(data => {
                document.getElementById('id').value = data.id;
                document.getElementById('label').value = data.label;
                document.getElementById('type').value = data.type;
                document.getElementById('maxLength').value = data.maxLength;
                document.getElementById('required').value = data.required;
                document.getElementById('options').value = data.options;
                document.getElementById('handler').value = data.handler;
            });
    }

    function deleteItem(itemId) {
        if (confirm('本当に削除しますか？')) {
            fetch(`/components/${id}/items/${itemId}`, { method: 'DELETE' })
                .then(response => response.json())
                .then(data => {
                    if (data.success) {
                        location.reload();
                    } else {
                        alert('削除に失敗しました。');
                    }
                });
        }
    }

    document.getElementById('componentForm').onsubmit = function(event) {
        event.preventDefault();
        const formData = new FormData(event.target);
        const jsonData = {};
        formData.forEach((value, key) => jsonData[key] = value);
        fetch(`/components/${id}`, {
            method: 'PUT',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify(jsonData)
        }).then(response => response.json())
        .then(data => {
            if (data.success) {
                alert('保存に成功しました。');
                location.reload();
            } else {


                alert('保存に失敗しました。');
            }
        });
    }
</script>
</body>
</html>
```

### サーバー側の例 (Golang)

```go
package main

import (
    "database/sql"
    "encoding/json"
    "log"
    "net/http"
    "strconv"

    "github.com/gin-gonic/gin"
    _ "github.com/mattn/go-sqlite3"
)

type Component struct {
    ID        string `json:"id"`
    Label     string `json:"label"`
    Type      string `json:"type"`
    MaxLength int    `json:"maxLength,omitempty"`
    Required  bool   `json:"required"`
    Options   string `json:"options,omitempty"`
    Handler   string `json:"handler"`
}

var db *sql.DB

func main() {
    var err error
    db, err = sql.Open("sqlite3", "./components.db")
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()

    createTable()

    router := gin.Default()
    router.LoadHTMLGlob("templates/*")

    router.GET("/components", getComponents)
    router.GET("/components/:id", getComponent)
    router.GET("/components/:id/items", getComponentItems)
    router.GET("/components/:id/items/:itemId", getComponentItem)
    router.POST("/components", createComponent)
    router.PUT("/components/:id", updateComponent)
    router.DELETE("/components/:id/items/:itemId", deleteComponentItem)

    log.Fatal(router.Run(":8080"))
}

func createTable() {
    sqlStmt := `
    CREATE TABLE IF NOT EXISTS components (
        id TEXT PRIMARY KEY,
        label TEXT,
        type TEXT,
        maxLength INTEGER,
        required BOOLEAN,
        options TEXT,
        handler TEXT
    );
    CREATE TABLE IF NOT EXISTS component_items (
        id INTEGER PRIMARY KEY AUTOINCREMENT,
        component_id TEXT,
        label TEXT,
        type TEXT,
        maxLength INTEGER,
        required BOOLEAN,
        options TEXT,
        handler TEXT,
        FOREIGN KEY(component_id) REFERENCES components(id)
    );
    `
    _, err := db.Exec(sqlStmt)
    if err != nil {
        log.Fatalf("%q: %s\n", err, sqlStmt)
        return
    }
}

func getComponents(c *gin.Context) {
    rows, err := db.Query("SELECT id, label FROM components")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var components []Component
    for rows.Next() {
        var component Component
        rows.Scan(&component.ID, &component.Label)
        components = append(components, component)
    }
    c.JSON(http.StatusOK, components)
}

func getComponent(c *gin.Context) {
    id := c.Param("id")
    row := db.QueryRow("SELECT id, label, type, maxLength, required, options, handler FROM components WHERE id = ?", id)
    var component Component
    err := row.Scan(&component.ID, &component.Label, &component.Type, &component.MaxLength, &component.Required, &component.Options, &component.Handler)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Component not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, component)
}

func getComponentItems(c *gin.Context) {
    id := c.Param("id")
    rows, err := db.Query("SELECT id, label, type, maxLength, required, options, handler FROM component_items WHERE component_id = ?", id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer rows.Close()

    var items []Component
    for rows.Next() {
        var item Component
        rows.Scan(&item.ID, &item.Label, &item.Type, &item.MaxLength, &item.Required, &item.Options, &item.Handler)
        items = append(items, item)
    }
    c.JSON(http.StatusOK, items)
}

func getComponentItem(c *gin.Context) {
    id := c.Param("itemId")
    row := db.QueryRow("SELECT id, label, type, maxLength, required, options, handler FROM component_items WHERE id = ?", id)
    var item Component
    err := row.Scan(&item.ID, &item.Label, &item.Type, &item.MaxLength, &item.Required, &item.Options, &item.Handler)
    if err != nil {
        if err == sql.ErrNoRows {
            c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
        } else {
            c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        }
        return
    }
    c.JSON(http.StatusOK, item)
}

func createComponent(c *gin.Context) {
    var newComponent Component
    if err := c.ShouldBindJSON(&newComponent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    stmt, err := db.Prepare("INSERT INTO components(id, label, type, maxLength, required, options, handler) VALUES(?, ?, ?, ?, ?, ?, ?)")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec(newComponent.ID, newComponent.Label, newComponent.Type, newComponent.MaxLength, newComponent.Required, newComponent.Options, newComponent.Handler)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, newComponent)
}

func updateComponent(c *gin.Context) {
    id := c.Param("id")
    var updatedComponent Component
    if err := c.ShouldBindJSON(&updatedComponent); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    stmt, err := db.Prepare("UPDATE components SET label = ?, type = ?, maxLength = ?, required = ?, options = ?, handler = ? WHERE id = ?")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec(updatedComponent.Label, updatedComponent.Type, updatedComponent.MaxLength, updatedComponent.Required, updatedComponent.Options, updatedComponent.Handler, id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, updatedComponent)
}

func deleteComponentItem(c *gin.Context) {
    id := c.Param("itemId")
    stmt, err := db.Prepare("DELETE FROM component_items WHERE id = ?")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    defer stmt.Close()
    _, err = stmt.Exec(id)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"success": true})
}
```