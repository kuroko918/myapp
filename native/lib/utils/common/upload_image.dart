import 'dart:io';
import 'package:firebase_storage/firebase_storage.dart';

Future<String> uploadImage(File image, String userId) async {
  final storage = FirebaseStorage.instance;
  final ref = storage.ref().child('images/users/$userId/thumbnail');
  final imageSnap = await ref.putFile(image);
  final imageURL = await imageSnap.ref.getDownloadURL();
  return imageURL;
}
