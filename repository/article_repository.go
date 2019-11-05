package repository

import (
	// model パッケージで定義した Article 構造体を利用するため、パッケージをインポート
	"go-tech-blog/model"
)

// ArticleList ...
/* 戻り値ありの関数 */
func ArticleList() ([]*model.Article, error) {
	// 実行する SQL
	query := `SELECT * FROM articles;`

	// データベースから取得した値を格納する変数を宣言
	var articles []*model.Article

	// Query を実行して、取得した値を変数に格納
	if err := db.Select(&articles, query); err != nil {
		return nil, err
	}

	return articles, nil
}
