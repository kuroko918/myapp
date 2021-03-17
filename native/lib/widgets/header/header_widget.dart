import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:myapp/utils/common/dialog.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';

class Header extends HookWidget implements PreferredSizeWidget {
  Header(this.title);

  final title;

  @override
  Size get preferredSize => const Size.fromHeight(60);

  @override
  Widget build(context) {
    return AppBar(
      title: Text(title),
      actions: [
        IconButton(
          icon: Icon(Icons.logout),
          tooltip: 'ログアウト',
          onPressed: () async {
            final result = await showAlertDialog(context, 'ログアウトしますか？', ['ログアウト', 'キャンセル']);
            if (!result) return;

            signOut();
            Navigator.pushReplacementNamed(context, '/login');
          },
        ),
      ],
    );
  }
}
