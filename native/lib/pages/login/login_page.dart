import 'package:flutter/material.dart';
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';

class LoginPage extends StatelessWidget {
  @override
  Widget build(context) {
    return Scaffold(
      appBar: AppBar(
        title: Text('Login'),
      ),
      body: Center(
        child: Builder(
          builder: (context) => ElevatedButton.icon(
            icon: Icon(
              Icons.person,
              color: Colors.white,
            ),
            label: Text(
              'GitHubでログインする',
              style: TextStyle(color: Colors.white),
            ),
            style: ElevatedButton.styleFrom(
              primary: Colors.black,
              shape: RoundedRectangleBorder(
                borderRadius: BorderRadius.circular(10),
              ),
            ),
            onPressed: () async {
              if (EasyLoading.isShow) return;

              try {
                EasyLoading.show(status: 'loading...');
                await signIn(context);
                EasyLoading.dismiss();
                Navigator.pushReplacementNamed(context, '/home');
              } catch (_) {
                EasyLoading.showError('ログインをキャンセルしました。');
              }
            },
          ),
        ),
      ),
    );
  }
}
