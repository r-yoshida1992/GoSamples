$(document).ready(function() {
	// ハンバーガーメニューの設定
	$('.sidenav').sidenav({
		edge: "right"
	});
	
	// codeブロックはpreタグを使用している為、自動で入る改行を削除する。
	$('code').text($('code').text().substr(1));

	// モーダルの処理
	$('.modal').modal();
})
