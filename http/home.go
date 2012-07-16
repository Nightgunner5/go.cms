package http

import (
	. "llamaslayers.net/go.cms/document"
	"net/http"
	"time"
)

func homeHandler(w http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "/" {
		notFoundHandler(w, req)
		return
	}
	DoPage(w, req, HomeDocument(), http.StatusOK)
}

func HomeDocument() *Document {
	return &Document{
		"Home",
		Content{
			&Article{
				"Cras porttitor rutrum metus pellentesque",
				time.Date(1004, time.March, 15, 5, 27, 0, 0, time.UTC),
				Content{
					&Paragraph{
						Content{
							&LeafElement{"Lorem ipsum dolor sit amet, consectetur adipiscing elit. Suspendisse vel augue nisi, malesuada rutrum eros. Vestibulum suscipit, augue nec posuere tempor, eros lectus congue nibh, nec pellentesque quam risus quis lacus. Aliquam erat volutpat. Integer bibendum arcu ac dui facilisis sit amet vulputate orci ornare. Fusce arcu ante, euismod id ornare quis, tempus sed nulla. Sed nibh mauris, lobortis sed elementum et, tristique vel nisi. Phasellus porta volutpat massa id aliquet. Vestibulum lacus dolor, fringilla sit amet consequat gravida, rutrum ac lectus. Maecenas vitae fermentum risus. Suspendisse vel augue elit, ac faucibus turpis. Nulla augue libero, gravida non bibendum vel, scelerisque et turpis. Aliquam odio ligula, interdum a rhoncus convallis, tristique eu magna."},
						},
					},
					&Paragraph{
						Content{
							&LeafElement{"Sed lorem urna, sodales faucibus commodo ut, rutrum aliquam orci. Integer accumsan quam eu enim vulputate elementum. Vestibulum sit amet risus orci. Fusce vestibulum faucibus sem ultricies pellentesque. Pellentesque nibh urna, sollicitudin sed feugiat facilisis, venenatis in nisl. Duis gravida, neque non euismod feugiat, sapien arcu sagittis urna, eget varius diam diam sed neque. Vestibulum mattis tellus ac libero porttitor in mattis elit porta. Fusce ac neque rhoncus metus volutpat facilisis id vitae leo. In at libero ut arcu iaculis euismod id id urna. In gravida, turpis pellentesque rhoncus eleifend, augue justo adipiscing leo, a semper neque lorem non justo. Mauris elementum dapibus lectus id dignissim."},
						},
					},
					&Paragraph{
						Content{
							&LeafElement{"Nullam non mi mauris, sit amet hendrerit risus. Sed eget eros sed velit tincidunt ultricies. Duis fringilla velit quis orci gravida et bibendum elit tempus. Pellentesque et lorem sit amet mauris sollicitudin fringilla. Maecenas diam erat, pharetra accumsan ornare ut, iaculis quis augue. Etiam iaculis, quam sit amet vehicula imperdiet, lectus massa ultrices leo, ut hendrerit augue urna at arcu. Nunc gravida molestie metus, sit amet tempor arcu bibendum sit amet. Nunc vestibulum pellentesque bibendum. Cras at sapien metus. Nunc iaculis, velit ut ullamcorper tincidunt, massa lectus pharetra lorem, sed pellentesque enim tellus lobortis lectus. Etiam ut enim libero, vitae varius mauris. Suspendisse id purus eget tellus fermentum mattis. Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
						},
					},
				},
			},
		},
	}
}
