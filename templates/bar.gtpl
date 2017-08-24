{{define "bar"}}
<style>
	nav a{
		width:150px;
		float:left;
		text-decoration: none;
	}
</style>

<nav>
	<a href="/profile" style="display: block; border:1px solid black;">
		Профиль
	</a>
	<a href="/friends" style="display: block; border:1px solid black;">
		Друзья
	</a>
	<a href="/messages" style="display: block; border:1px solid black;">
		Сообщения
	</a>
	<a href="/photos" style="display: block; border:1px solid black;">
		Фотографии
	</a>
	<a href="/photos" style="display: block; border:1px solid black;">
		Видеозаписи
	</a>
	
	<a href="/photos" style="display: block; border:1px solid black;">
		Музыка
	</a>

	<a href="/photos" style="display: block; border:1px solid black;">
		Выйти
	</a>
	
	
</nav>
{{end}}
