import 'package:firebase_core/firebase_core.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart' as DotEnv;
import 'package:flutter_easyloading/flutter_easyloading.dart';
import 'package:flutter/material.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:myapp/config/flutter_easyloading/config_loading.dart';
import 'package:myapp/config/logger/logger.dart';
import 'package:myapp/pages/home/home_page.dart';
import 'package:myapp/pages/login/login_page.dart';
import 'package:myapp/utils/firebase/firebase_auth.dart';

void main() async {
  WidgetsFlutterBinding.ensureInitialized();
  await Firebase.initializeApp();
  await DotEnv.load();
  runApp(
    ProviderScope(
      observers: [Logger()],
      child: MyApp(),
    ),
  );
  configLoading();
}

class MyApp extends StatelessWidget {
  @override
  Widget build(context) {
    return MaterialApp(
      initialRoute: isSignedIn() ? '/home' : '/login',
      builder: EasyLoading.init(),
      routes: {
        '/home': (_) => HomePage(),
        '/login': (_) => LoginPage(),
      },
    );
  }
}
