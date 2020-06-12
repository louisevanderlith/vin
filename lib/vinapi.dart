import 'dart:async';
import 'dart:html';

import '../pathlookup.dart';
import 'requester.dart';

Future<HttpRequest> validateVIN(String vin) async {
  final url = await buildPath("VIN.API", "validate", [vin]);

  return invokeService("GET", url, true, "");
}
