import 'package:hooks_riverpod/hooks_riverpod.dart';

class HomeState extends StateNotifier<int> {
  HomeState() : super(0);

  int currentIndex = 0;

  void changeIndex(int index) {
    currentIndex = index;
    state = currentIndex;
  }
}
