syntax = "proto3";

package environmental_alert.v1;

option go_package = "github.com/anhvanhoa/sf-proto/gen/environmental_alert/v1;proto_environmental_alert";

import "google/protobuf/timestamp.proto";
import "google/protobuf/empty.proto";
import "buf/validate/validate.proto";
import "common/v1/common.proto";

service EnvironmentalAlertService {
  rpc CreateEnvironmentalAlert(CreateEnvironmentalAlertRequest)
      returns (CreateEnvironmentalAlertResponse);
  rpc GetEnvironmentalAlert(GetEnvironmentalAlertRequest)
      returns (GetEnvironmentalAlertResponse);
  rpc UpdateEnvironmentalAlert(UpdateEnvironmentalAlertRequest)
      returns (UpdateEnvironmentalAlertResponse);
  rpc DeleteEnvironmentalAlert(DeleteEnvironmentalAlertRequest)
      returns (DeleteEnvironmentalAlertResponse);
  rpc ListEnvironmentalAlerts(ListEnvironmentalAlertsRequest)
      returns (ListEnvironmentalAlertsResponse);
  rpc GetEnvironmentalAlertStatistics(google.protobuf.Empty)
      returns (GetEnvironmentalAlertStatisticsResponse);
}

// Main Message
message EnvironmentalAlert {
  string id = 1 [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string device_id = 2
      [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string alert_type = 3 [
    (buf.validate.field).string.in = "temperature_high",
    (buf.validate.field).string.in = "temperature_low",
    (buf.validate.field).string.in = "humidity_high",
    (buf.validate.field).string.in = "humidity_low",
    (buf.validate.field).string.in = "ph_high",
    (buf.validate.field).string.in = "ph_low",
    (buf.validate.field).string.in = "water_shortage",
    (buf.validate.field).string.in = "equipment_failure"
  ];
  double current_value = 4
      [ (buf.validate.field).double = {gte : -1000, lte : 1000} ];
  double threshold_value = 5
      [ (buf.validate.field).double = {gte : -1000, lte : 1000} ];
  string threshold_type = 6 [
    (buf.validate.field).string.in = "min",
    (buf.validate.field).string.in = "max",
    (buf.validate.field).string.in = "range"
  ];
  string severity = 7 [
    (buf.validate.field).string.in = "info",
    (buf.validate.field).string.in = "warning",
    (buf.validate.field).string.in = "critical",
    (buf.validate.field).string.in = "emergency"
  ];
  int32 priority = 8 [ (buf.validate.field).int32 = {gte : 1, lte : 10} ];
  google.protobuf.Timestamp triggered_at = 9;
  google.protobuf.Timestamp acknowledged_at = 10;
  string acknowledged_by = 11 [ (buf.validate.field).string = {max_len : 100} ];
  google.protobuf.Timestamp resolved_at = 12;
  string resolved_by = 13 [ (buf.validate.field).string = {max_len : 100} ];
  bool auto_resolve = 14;
  string status = 15 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "acknowledged",
    (buf.validate.field).string.in = "resolved",
    (buf.validate.field).string.in = "escalated",
    (buf.validate.field).string.in = "ignored"
  ];
  string auto_action_taken = 16
      [ (buf.validate.field).string = {max_len : 500} ];
  int32 escalation_level = 17
      [ (buf.validate.field).int32 = {gte : 0, lte : 5} ];
  string impact_assessment = 18
      [ (buf.validate.field).string = {max_len : 1000} ];
  string resolution_notes = 19
      [ (buf.validate.field).string = {max_len : 1000} ];
  string created_by = 20
      [ (buf.validate.field).string = {min_len : 1, max_len : 100} ];
  google.protobuf.Timestamp created_at = 21;
  google.protobuf.Timestamp updated_at = 22;
}

message AlertFilters {
  string device_id = 1 [ (buf.validate.field).string = {max_len : 50} ];
  string alert_type = 2 [
    (buf.validate.field).string.in = "temperature_high",
    (buf.validate.field).string.in = "temperature_low",
    (buf.validate.field).string.in = "humidity_high",
    (buf.validate.field).string.in = "humidity_low",
    (buf.validate.field).string.in = "ph_high",
    (buf.validate.field).string.in = "ph_low",
    (buf.validate.field).string.in = "water_shortage",
    (buf.validate.field).string.in = "equipment_failure"
  ];
  string status = 3 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "acknowledged",
    (buf.validate.field).string.in = "resolved",
    (buf.validate.field).string.in = "escalated",
    (buf.validate.field).string.in = "ignored"
  ];
  string severity = 4 [
    (buf.validate.field).string.in = "info",
    (buf.validate.field).string.in = "warning",
    (buf.validate.field).string.in = "critical",
    (buf.validate.field).string.in = "emergency"
  ];
  int32 priority = 5 [ (buf.validate.field).int32 = {gte : 1, lte : 10} ];
  google.protobuf.Timestamp start_date = 6;
  google.protobuf.Timestamp end_date = 7;
  bool is_overdue = 8;
  int32 escalation_level = 9
      [ (buf.validate.field).int32 = {gte : 0, lte : 5} ];
  bool auto_resolve = 10;
  string created_by = 11 [ (buf.validate.field).string = {max_len : 100} ];
  string acknowledged_by = 12 [ (buf.validate.field).string = {max_len : 100} ];
  string resolved_by = 13 [ (buf.validate.field).string = {max_len : 100} ];
}

message AlertStatistics {
  int64 total_alerts = 1;
  int64 active_alerts = 2;
  int64 resolved_alerts = 3;
  int64 overdue_alerts = 4;
}

message CreateEnvironmentalAlertRequest {
  string device_id = 1
      [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string alert_type = 2 [
    (buf.validate.field).string.in = "temperature_high",
    (buf.validate.field).string.in = "temperature_low",
    (buf.validate.field).string.in = "humidity_high",
    (buf.validate.field).string.in = "humidity_low",
    (buf.validate.field).string.in = "ph_high",
    (buf.validate.field).string.in = "ph_low",
    (buf.validate.field).string.in = "water_shortage",
    (buf.validate.field).string.in = "equipment_failure"
  ];
  double current_value = 3
      [ (buf.validate.field).double = {gte : -1000, lte : 1000} ];
  double threshold_value = 4
      [ (buf.validate.field).double = {gte : -1000, lte : 1000} ];
  string threshold_type = 5 [
    (buf.validate.field).string.in = "min",
    (buf.validate.field).string.in = "max",
    (buf.validate.field).string.in = "range"
  ];
  string severity = 6 [
    (buf.validate.field).string.in = "info",
    (buf.validate.field).string.in = "warning",
    (buf.validate.field).string.in = "critical",
    (buf.validate.field).string.in = "emergency"
  ];
  int32 priority = 7 [ (buf.validate.field).int32 = {gte : 1, lte : 10} ];
  bool auto_resolve = 8;
  string auto_action_taken = 9
      [ (buf.validate.field).string = {max_len : 500} ];
  int32 escalation_level = 10
      [ (buf.validate.field).int32 = {gte : 0, lte : 5} ];
  string impact_assessment = 11
      [ (buf.validate.field).string = {max_len : 1000} ];
  string created_by = 12
      [ (buf.validate.field).string = {min_len : 1, max_len : 100} ];
}

message CreateEnvironmentalAlertResponse { EnvironmentalAlert alert = 1; }

message GetEnvironmentalAlertRequest {
  string id = 1 [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
}

message GetEnvironmentalAlertResponse { EnvironmentalAlert alert = 1; }

message UpdateEnvironmentalAlertRequest {
  string id = 1 [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string device_id = 2
      [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
  string alert_type = 3 [
    (buf.validate.field).string.in = "temperature_high",
    (buf.validate.field).string.in = "temperature_low",
    (buf.validate.field).string.in = "humidity_high",
    (buf.validate.field).string.in = "humidity_low",
    (buf.validate.field).string.in = "ph_high",
    (buf.validate.field).string.in = "ph_low",
    (buf.validate.field).string.in = "water_shortage",
    (buf.validate.field).string.in = "equipment_failure"
  ];
  double current_value = 4
      [ (buf.validate.field).double = {gte : -1000, lte : 1000} ];
  double threshold_value = 5
      [ (buf.validate.field).double = {gte : -1000, lte : 1000} ];
  string threshold_type = 6 [
    (buf.validate.field).string.in = "min",
    (buf.validate.field).string.in = "max",
    (buf.validate.field).string.in = "range"
  ];
  string severity = 7 [
    (buf.validate.field).string.in = "info",
    (buf.validate.field).string.in = "warning",
    (buf.validate.field).string.in = "critical",
    (buf.validate.field).string.in = "emergency"
  ];
  int32 priority = 8 [ (buf.validate.field).int32 = {gte : 1, lte : 10} ];
  string status = 9 [
    (buf.validate.field).string.in = "active",
    (buf.validate.field).string.in = "acknowledged",
    (buf.validate.field).string.in = "resolved",
    (buf.validate.field).string.in = "escalated",
    (buf.validate.field).string.in = "ignored"
  ];
  bool auto_resolve = 10;
  string auto_action_taken = 11
      [ (buf.validate.field).string = {max_len : 500} ];
  int32 escalation_level = 12
      [ (buf.validate.field).int32 = {gte : 0, lte : 5} ];
  string impact_assessment = 13
      [ (buf.validate.field).string = {max_len : 1000} ];
  string resolution_notes = 14
      [ (buf.validate.field).string = {max_len : 1000} ];
  string updated_by = 15
      [ (buf.validate.field).string = {min_len : 1, max_len : 100} ];
}

message UpdateEnvironmentalAlertResponse { EnvironmentalAlert alert = 1; }

message DeleteEnvironmentalAlertRequest {
  string id = 1 [ (buf.validate.field).string = {min_len : 1, max_len : 50} ];
}

message DeleteEnvironmentalAlertResponse {
  bool success = 1;
  string message = 2;
}

message ListEnvironmentalAlertsRequest {
  common.PaginationRequest pagination = 1;
  AlertFilters filters = 2;
}

message ListEnvironmentalAlertsResponse {
  repeated EnvironmentalAlert alerts = 1;
  common.PaginationResponse pagination = 2;
}

message GetEnvironmentalAlertStatisticsResponse {
  AlertStatistics statistics = 1;
}