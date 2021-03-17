import 'package:cloud_firestore/cloud_firestore.dart';
import 'package:firebase_auth/firebase_auth.dart';

Future<bool> isUserExisting(User user) async {
  final userSnap = await FirebaseFirestore.instance.collection('users').doc(user.uid).get();
  return userSnap.exists;
}

void storeUser(User user) {
  FirebaseFirestore.instance.collection('users').doc(user.uid).set({
    'id': user.uid,
    'name': user.displayName,
    'email': user.email,
    'avatar': user.photoURL,
    'createdAt': FieldValue.serverTimestamp(),
    'updatedAt': FieldValue.serverTimestamp(),
  });
}
