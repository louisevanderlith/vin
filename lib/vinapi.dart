import 'dart:async';
import 'dart:html';

import 'package:mango_ui/requester.dart';

Future<HttpRequest> validateVIN(String vin) async {
  var apiroute = getEndpoint("vin");
  var url = "${apiroute}/validate/${vin}";

  return invokeService("GET", url, null);
}
