import 'dart:io';
import 'package:image_picker/image_picker.dart';

Future<File> getImageFromCamera() async {
  final picker = ImagePicker();
  final pickedFile = await picker.getImage(source: ImageSource.camera);

  print("pickedFile");

  print(pickedFile);

  if (pickedFile == null) return null;

  return File(pickedFile.path);
}

Future<File> getImageFromGallery() async {
  final picker = ImagePicker();
  final pickedFile = await picker.getImage(source: ImageSource.gallery);

  if (pickedFile == null) return null;

  return File(pickedFile.path);
}
