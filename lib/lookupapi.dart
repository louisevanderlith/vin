import 'dart:html';

import 'package:mango_ui/requester.dart';

Future<HttpRequest> fetchManufacturers(String year) async {
  var apiroute = getEndpoint("vin");
  var url = "${apiroute}/lookup/manufacturers/${year}";

  return invokeService("GET", url, null);
}

Future<HttpRequest> fetchModels(String year, String manufacturer) async {
  var apiroute = getEndpoint("vin");
  var url = "${apiroute}/lookup/models/${year}/${manufacturer}";

  return invokeService("GET", url, null);
}

Future<HttpRequest> fetchTrims(
    String year, String manufacturer, String model) async {
  var apiroute = getEndpoint("vin");
  var url = "${apiroute}/lookup/trim/${year}/${manufacturer}/${model}";

  return invokeService("GET", url, null);
}
