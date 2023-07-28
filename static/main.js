const topImage = document.querySelector("topImage");
const cl = topImage.classList;
cl.remove("low");
cl.add("high");

fetch("/static/bg_1.jpg")
  .then(function (res) {
    return res.blob();
  })
  .then(function (blob) {
    url = URL.createObjectURL(blob);

    const img = new Image();
    img.src = url;

    const body = document.querySelector("body");
    body.appendChild(img);
  });
