#= require jquery

$ ->

  isBlank = (str) ->
    if str?
      return str.trim() is ""
    return true

  displayUser = (user) ->
    $("table tbody").append("""
      <tr>
        <td>#{user.id}</td>
        <td>#{user.name}</td>
        <td>#{user.email}</td>
      </tr>
    """)
    $("#users").show()

  db = openDatabase('metacasts', '1.0', 'a simple db', 5 * 1024 * 1024)

  db.transaction (tx) ->
    tx.executeSql('''
      CREATE TABLE "users" (
        "id" integer NOT NULL PRIMARY KEY AUTOINCREMENT,
        "name" text NOT NULL,
        "email" text NOT NULL
      );
    ''')

  db.select = (stmt, values, callback) ->
    @transaction (tx) ->
      tx.executeSql stmt, (values || []), (tx, results) ->
        len = results.rows.length
        ind = 0
        until ind is len
          row = results.rows.item(ind)
          displayUser(row)
          if callback?
            callback(row)
          ind++

  $("#user-form").submit (e) ->
    e?.preventDefault()
    name = $("#user-name").val()
    email = $("#user-email").val()
    if isBlank(email) or isBlank(name)
      alert "Please fill in the whole form!"
    else
      db.transaction (tx) ->
        tx.executeSql 'INSERT INTO users (name, email) VALUES (?, ?)', [name, email], (tx, results) ->
          db.select("SELECT * FROM users WHERE id = ?", [results.insertId])

  db.select('SELECT * FROM users;')