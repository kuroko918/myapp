import 'package:flutter/material.dart';
import 'package:flutter_hooks/flutter_hooks.dart';

class Footer extends HookWidget {
  Footer(this.currentIndex, this.changeIndex);

  final int currentIndex;
  final void Function(int) changeIndex;

  @override
  Widget build(context) {
    return BottomNavigationBar(
      items: [
        BottomNavigationBarItem(
          icon: Icon(Icons.chat),
          label: 'chat',
        ),
        BottomNavigationBarItem(
          icon: Icon(Icons.account_circle),
          label: 'profile',
        ),
      ],
      currentIndex: currentIndex,
      onTap: (index) => changeIndex(index),
    );
  }
}
