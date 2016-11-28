<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
  <title>ExpSender</title>
  <link rel="stylesheet" type="text/css" href="static/css/text.css">
</head>
<body>
	<form>
	<input type="text" name="apiKey" placeholder="apiKey">
	<input type="submit" value="Получить листы">
	<p>{{.Code1}}
	</form>
	<form method="POST">
	<input type="text" name="apiKey" placeholder="apiKey" required> <br>
	<input type="text" name="SeedList" placeholder="SeedList" required> <br>
	<input type="text" name="FromName" placeholder="FromName" required> <br> 
	<input type="text" name="FromEmail" placeholder="FromEmail" required> <br>
	<input type="text" name="Subject" placeholder="Subject" required> <br>
	<input type="submit" value="Отправить листы"><br>
	<p>{{.Code2}}
	</form>
<table align="left" border="1" width="500" cellspacing="0" cellpadding="10">
<caption>{{.Table1}}</caption>
<tr><th>Id</th><th>Name</th><th>FriendlyName</th><th>Language</th></tr>
{{.SingleOptIn}}
</table>
<table align="right" border="1" width="500" cellspacing="0" cellpadding="10">
<caption>{{.Table2}}</caption>
<tr><th>Id</th><th>Name</th><th>FriendlyName</th><th>Language</th></tr>
{{.DoubleOptIn}}
</table>
</body>
</html>
