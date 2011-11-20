var svgns = "http://www.w3.org/2000/svg";
var SIZE = 10;

function blah() {
	console.log("blah!");
	svg = document.getElementById("svg");
	
	rects(svg, water, {fill: "blue", stroke: "midnightblue"});
	rects(svg, path, {fill:"red", stroke: "red", opacity: 0.5});
}

function rects(parent, points, attrs) {
	for (var j=0; j< points.length; j++ ) {
		var r = document.createElementNS(svgns, "rect");
		var x = points[j][1];
		var y = points[j][0];
		r.setAttribute("x", x * SIZE);
		r.setAttribute("y", y * SIZE);
		r.setAttribute("width", SIZE);
		r.setAttribute("height", SIZE);

		for (a in attrs) {
			r.setAttribute(a, attrs[a]);
		}

		parent.appendChild(r)
	}
}
