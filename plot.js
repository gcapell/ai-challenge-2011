var svgns = "http://www.w3.org/2000/svg";
var SIZE = 10;

function hello(j) {
	if (LOC == undefined) {
		LOC = document.getElementById("location");
	}
	LOC.innerHTML = "" + popped[j];
}

function blah() {
	console.log("blah!");
	svg = document.getElementById("svg");
	
	rects(svg, water, {fill: "blue", stroke: "midnightblue", });
	rects(svg, path, {fill:"red", stroke: "red", opacity: 0.5});
	//rects(svg, path2, {fill:"green", stroke: "red", opacity: 0.3});
	// rects(svg, expanded, {fill:"orange", stroke: "red", opacity: 0.3});
	rects(svg, popped, {fill:"pink", stroke: "red", opacity: 0.2}, true);
}

var LOC;

function rects(parent, points, attrs, callback) {
	for (var j=0; j< points.length; j++ ) {
		var r = document.createElementNS(svgns, "rect");
		var x = points[j][1];
		var y = points[j][0];
		r.setAttribute("x", x * SIZE);
		r.setAttribute("y", y * SIZE);
		r.setAttribute("width", SIZE);
		r.setAttribute("height", SIZE);
		if (callback) {
			r.setAttribute("onmouseover", "hello(" + j + ")"  );
		}

		
		for (a in attrs) {
			r.setAttribute(a, attrs[a]);
		}

		parent.appendChild(r)
	}
}
