package main

import (
	"fmt"
	"net/http"
)

const tHome = `
<div class="container">
	<div class="col-xs-12">
		<div class="jumbotron">
			<h1>Byte-order mark remover</h1>
			<p>Remove all unneeded byte-order marks from the top of your JSON files.</p>
		</div>
	</div>
</div>

<div class="container">
	<div class="col-xs-12 col-md-8 col-md-offset-2">
		<form class="form-horizontal" enctype="multipart/form-data" method="post" action="/upload">
			<div class="form-group">
				<p>This website takes a JSON file and parses it to remove the JSON byte-order marks that some file editors may left behind.</p>
			</div>
			<div class="form-group">
				<label class="col-sm-2 control-label">File:</label>
				<div class="col-sm-10">
					<input type="file" class="form-control" name="bomfile" accept=".json">
				</div>
			</div>
			<div class="form-group">
				<div class="col-sm-offset-2 col-sm-10">
					<button type="submit" class="btn btn-default">Upload!</button>
				</div>
			</div>
		</form>
	</div>
</div>
`

func root(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	}

	fmt.Fprint(w, getTemplate(tHome))
}
