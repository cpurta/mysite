var svg = d3.select("svg"),
margin = {top: 50, right: 20, bottom: 30, left: 50},
width = +svg.attr("width") - margin.left - margin.right,
height = +svg.attr("height") - margin.top - margin.bottom,
g = svg.append("g").attr("transform", "translate(" + margin.left + "," + margin.top + ")");

var strictIsoParse = d3.utcParse("%Y-%m-%dT%H:%M:%S%Z");

var x = d3.scaleTime().range([0, width]);
var y = d3.scaleLinear().range([height, 0]);
var z = d3.scaleOrdinal(d3.schemeCategory10);

var line = d3.line()
    .x(function(d) { return x(d.date); })
    .y(function(d) { return y(d.price); });

d3.csv("/public/data/data.csv", function(d) {
  d.date = strictIsoParse(d.date);
  d.price = +d.price;
  return d;
}, function(error, data) {
  if (error) throw error;

  x.domain(d3.extent(data, function(d) { return d.date; }));
  y.domain(d3.extent(data, function(d) { return d.price; }));

  g.append("g")
      .attr("transform", "translate(0," + height + ")")
      .call(d3.axisBottom(x))
    .select(".domain")
      .remove();

  g.append("g")
      .call(d3.axisLeft(y))
    .append("text")
      .attr("fill", "#000")
      .attr("transform", "rotate(-90)")
      .attr("y", 6)
      .attr("dy", "0.71em")
      .attr("text-anchor", "end")
      .text("Price (USD)");

  g.append("path")
      .datum(data)
      .attr("fill", "none")
      .attr("stroke", "steelblue")
      .attr("stroke-linejoin", "round")
      .attr("stroke-linecap", "round")
      .attr("stroke-width", 2.0)
      .attr("d", line);
});
