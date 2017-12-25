<!DOCTYPE html>

<html>
<head>
  <title>KDLUCK</title>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">

  <style type="text/css">
    *,body {
      margin: 0px;
      padding: 0px;
    }

    body {
      margin: 0px;
      font-family: "Helvetica Neue", Helvetica, Arial, sans-serif;
      font-size: 14px;
      line-height: 20px;
      background-color: #fff;
    }

    header,
    footer {
      width: 960px;
      margin-left: auto;
      margin-right: auto;
    }

    .logo {
      text-align: center;
      font-size: 42px;
      padding: 50px 0 50px;
      font-weight: normal;
      text-shadow: 0px 1px 2px #ddd;
    }

    header {
      padding: 0 0;
    }

    footer {
      line-height: 1.8;
      text-align: center;
      padding: 30px 0;
      color: #999;
    }

    .img {
      text-align: center;
    }

    a {
      color: #444;
      text-decoration: none;
    }

    .backdrop {
      position: absolute;
      width: 100%;
      height: 100%;
      box-shadow: inset 0px 0px 100px #ddd;
      z-index: -1;
      top: 0px;
      left: 0px;
    }
  </style>
</head>

<body>
  <header>
    <h1 class="logo">凯迪LUCK有钱有限公司</h1>
    <div class="img">
      <img src="{{.img}}" width="35%">
    </div>
  </header>
  <footer>
    <div>访问量  uv:{{.uv}} / pv:{{.pv}}</div>
  </footer>
  <div class="backdrop"></div>

  <script src="/static/js/reload.min.js"></script>
</body>
</html>
