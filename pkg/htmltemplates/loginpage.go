package htmltemplates

const Login = `
<!DOCTYPE HTML>
<html>
<title>LoginPage</title>
<body>
<form method="post">
  <input type="hidden" name="challenge" value="{{.Challenge}}">
  Username: <input type="text" name="username" value="{{.PrefillUser}}"/>
  <br/><br/>
  Password: <input type="password" name="password" value= "{{.PrefillPassword}}" />
  <br/><br/>
  <input type="submit" value="submit" />
</form>
<br/><br/>
{{ if .LoginFailed }}
Username and/or password incorrect
{{- end }}
{{ if .LoginError }}
An error occurred while authenticating the user.
{{- end }}
</body>
</html>
`

const Consent = `
<!DOCTYPE HTML>
<html>
<title>ConsentPage</title>
<body>

Hi {{.User}},<br/><br/>

Application <strong>{{.ClientId}}</strong> (Name: {{.ClientName}}) wants access resources on your behalf.<br/><br/>
The application requests access to the following items:

<form method="post">
  {{block "list" .Scope}}{{"<br/>"}}{{range .}}{{printf "<input type='checkbox' name='scope_%s' value='checked' checked> %s<br/>" . .}}{{end}}{{end}}
  <br/>
  <input type="hidden" name="challenge" value="{{.Challenge}}">
  <input type="submit" name="userconsent" value="accept" />
  <input type="submit" name="userconsent" value="reject" />
</form>

</body>
</html>
`
