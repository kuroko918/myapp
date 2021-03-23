import 'dart:io';
import 'package:flutter/material.dart';
import 'package:myapp/utils/common/get_image.dart';

Future<File> showGetImageModal(BuildContext context) {
  return showModalBottomSheet(
    context: context,
    builder: (context) => Column(
      mainAxisSize: MainAxisSize.min,
      children: [
        ListTile(
          leading: Icon(Icons.camera),
          title: Text('Camera'),
          onTap: () async {
            final image = await getImageFromCamera();
            Navigator.of(context).pop(image);
          },
        ),
        ListTile(
          leading: Icon(Icons.photo_album_outlined),
          title: Text('Carelly'),
          onTap: () async {
            final image = await getImageFromGallery();
            Navigator.of(context).pop(image);
          },
        ),
      ],
    ),
  );
}
