import 'package:flutter/material.dart';

Future<bool> showAlertDialog(BuildContext context, String title, List<String> texts) {
  return showDialog(
    context: context,
    barrierDismissible: false,
    builder: (context) => AlertDialog(
      title: Text('$title'),
      actions: [
        TextButton(
          child: Text('${texts[0]}'),
          onPressed: () => Navigator.of(context).pop(true),
        ),
        TextButton(
          child: Text('${texts[1]}'),
          onPressed: () => Navigator.of(context).pop(false),
        ),
      ],
    ),
  );
}
