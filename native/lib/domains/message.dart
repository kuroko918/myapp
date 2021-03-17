import 'package:flutter/foundation.dart';
import 'package:freezed_annotation/freezed_annotation.dart';
import 'package:myapp/domains/user.dart';

part 'message.freezed.dart';
part 'message.g.dart';

@freezed
abstract class Message with _$Message {
  const factory Message({
    String id,
    String content,
    String userId,
    DateTime createdAt,
    DateTime updatedAt,
    @JsonKey(name: 'user') User user,
  }) = _Message;

  factory Message.fromJson(Map<String, dynamic> json) => _$MessageFromJson(json);
}
