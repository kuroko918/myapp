import 'package:firebase_auth/firebase_auth.dart';
import 'package:flutter/material.dart';
import 'package:flutter_dotenv/flutter_dotenv.dart';
import 'package:github_sign_in/github_sign_in.dart';
import 'package:myapp/utils/firebase/firestore.dart';

User currentUser() => FirebaseAuth.instance.currentUser;

bool isSignedIn() => FirebaseAuth.instance.currentUser != null;

void signOut() => FirebaseAuth.instance.signOut();

Future<String> getIdToken() async {
  final user = FirebaseAuth.instance.currentUser;
  return await user.getIdToken();
}

Future<void> signIn(BuildContext context) async {
  final GitHubSignIn gitHubSignIn = GitHubSignIn(
    clientId: env['GITHUB_CLIENT_ID'],
    clientSecret: env['GITHUB_CLIENT_SECRET'],
    redirectUrl: env['GITHUB_REDIRECT_URL'],
  );

  final result = await gitHubSignIn.signIn(context);
  final githubAuthCredential = GithubAuthProvider.credential(result.token);
  final userCredential = await FirebaseAuth.instance.signInWithCredential(githubAuthCredential);
  final user = userCredential.user;

  if (await isUserExisting(user)) return;

  storeUser(user);
}
