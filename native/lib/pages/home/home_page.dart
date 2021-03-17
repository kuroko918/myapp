import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';
import 'package:hooks_riverpod/hooks_riverpod.dart';
import 'package:myapp/pages/chat/chat_page.dart';
import 'package:myapp/pages/home/home_model.dart';
import 'package:myapp/pages/profile/profile_page.dart';
import 'package:myapp/widgets/footer/footer_widget.dart';

final homeProvider = StateNotifierProvider((_) => HomeState());

class HomePage extends HookWidget {
  final List<Widget> tabs = [
    ChatPage(),
    ProfilePage(),
  ];

  @override
  Widget build(context) {
    final currentIndex = useProvider(homeProvider.state);

    return Scaffold(
      body: tabs[currentIndex],
      bottomNavigationBar: Footer(currentIndex, context.read(homeProvider).changeIndex),
    );
  }
}
