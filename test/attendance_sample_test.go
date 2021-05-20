// Code generated by lark_sdk_gen. DO NOT EDIT.

package test

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/chyroc/lark"
)

func Test_Attendance_Sample_Failed(t *testing.T) {
	as := assert.New(t)

	t.Run("request failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		cli.Mock().MockGetTenantAccessToken(mockGetTenantAccessTokenFailed)
		cli.Mock().MockGetAppAccessToken(mockGetTenantAccessTokenFailed)
		moduleCli := cli.Attendance

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateAttendanceUserSettings(ctx, &lark.UpdateAttendanceUserSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateAttendanceGroup(ctx, &lark.CreateUpdateAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceShiftByID(ctx, &lark.GetAttendanceShiftByIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceShiftByName(ctx, &lark.GetAttendanceShiftByNameReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceStatisticsData(ctx, &lark.GetAttendanceStatisticsDataReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceStatisticsHeader(ctx, &lark.GetAttendanceStatisticsHeaderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateAttendanceUserStatisticsSettings(ctx, &lark.UpdateAttendanceUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserStatisticsSettings(ctx, &lark.GetAttendanceUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateAttendanceUserDailyShift(ctx, &lark.CreateUpdateAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "failed")
		})
	})

	t.Run("request mock failed", func(t *testing.T) {
		cli := AppALLPermission.Ins()
		moduleCli := cli.Attendance

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceUpdateAttendanceUserSettings(func(ctx context.Context, request *lark.UpdateAttendanceUserSettingsReq, options ...lark.MethodOptionFunc) (*lark.UpdateAttendanceUserSettingsResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUpdateAttendanceUserSettings()

			_, _, err := moduleCli.UpdateAttendanceUserSettings(ctx, &lark.UpdateAttendanceUserSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceUploadAttendanceFile(func(ctx context.Context, request *lark.UploadAttendanceFileReq, options ...lark.MethodOptionFunc) (*lark.UploadAttendanceFileResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUploadAttendanceFile()

			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateUpdateAttendanceGroup(func(ctx context.Context, request *lark.CreateUpdateAttendanceGroupReq, options ...lark.MethodOptionFunc) (*lark.CreateUpdateAttendanceGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateUpdateAttendanceGroup()

			_, _, err := moduleCli.CreateUpdateAttendanceGroup(ctx, &lark.CreateUpdateAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceDeleteAttendanceGroup(func(ctx context.Context, request *lark.DeleteAttendanceGroupReq, options ...lark.MethodOptionFunc) (*lark.DeleteAttendanceGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceDeleteAttendanceGroup()

			_, _, err := moduleCli.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceGroup(func(ctx context.Context, request *lark.GetAttendanceGroupReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceGroupResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceGroup()

			_, _, err := moduleCli.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateAttendanceShift(func(ctx context.Context, request *lark.CreateAttendanceShiftReq, options ...lark.MethodOptionFunc) (*lark.CreateAttendanceShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateAttendanceShift()

			_, _, err := moduleCli.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceDeleteAttendanceShift(func(ctx context.Context, request *lark.DeleteAttendanceShiftReq, options ...lark.MethodOptionFunc) (*lark.DeleteAttendanceShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceDeleteAttendanceShift()

			_, _, err := moduleCli.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceShiftByID(func(ctx context.Context, request *lark.GetAttendanceShiftByIDReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceShiftByIDResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceShiftByID()

			_, _, err := moduleCli.GetAttendanceShiftByID(ctx, &lark.GetAttendanceShiftByIDReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceShiftByName(func(ctx context.Context, request *lark.GetAttendanceShiftByNameReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceShiftByNameResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceShiftByName()

			_, _, err := moduleCli.GetAttendanceShiftByName(ctx, &lark.GetAttendanceShiftByNameReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceStatisticsData(func(ctx context.Context, request *lark.GetAttendanceStatisticsDataReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceStatisticsDataResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceStatisticsData()

			_, _, err := moduleCli.GetAttendanceStatisticsData(ctx, &lark.GetAttendanceStatisticsDataReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceStatisticsHeader(func(ctx context.Context, request *lark.GetAttendanceStatisticsHeaderReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceStatisticsHeaderResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceStatisticsHeader()

			_, _, err := moduleCli.GetAttendanceStatisticsHeader(ctx, &lark.GetAttendanceStatisticsHeaderReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceUpdateAttendanceUserStatisticsSettings(func(ctx context.Context, request *lark.UpdateAttendanceUserStatisticsSettingsReq, options ...lark.MethodOptionFunc) (*lark.UpdateAttendanceUserStatisticsSettingsResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceUpdateAttendanceUserStatisticsSettings()

			_, _, err := moduleCli.UpdateAttendanceUserStatisticsSettings(ctx, &lark.UpdateAttendanceUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceUserStatisticsSettings(func(ctx context.Context, request *lark.GetAttendanceUserStatisticsSettingsReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserStatisticsSettingsResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserStatisticsSettings()

			_, _, err := moduleCli.GetAttendanceUserStatisticsSettings(ctx, &lark.GetAttendanceUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceUserDailyShift(func(ctx context.Context, request *lark.GetAttendanceUserDailyShiftReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserDailyShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserDailyShift()

			_, _, err := moduleCli.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceUserTask(func(ctx context.Context, request *lark.GetAttendanceUserTaskReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserTaskResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserTask()

			_, _, err := moduleCli.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceUserFlow(func(ctx context.Context, request *lark.GetAttendanceUserFlowReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserFlow()

			_, _, err := moduleCli.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceBatchGetAttendanceUserFlow(func(ctx context.Context, request *lark.BatchGetAttendanceUserFlowReq, options ...lark.MethodOptionFunc) (*lark.BatchGetAttendanceUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceBatchGetAttendanceUserFlow()

			_, _, err := moduleCli.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceBatchCreateAttendanceUserFlow(func(ctx context.Context, request *lark.BatchCreateAttendanceUserFlowReq, options ...lark.MethodOptionFunc) (*lark.BatchCreateAttendanceUserFlowResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceBatchCreateAttendanceUserFlow()

			_, _, err := moduleCli.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceUserTaskRemedy(func(ctx context.Context, request *lark.GetAttendanceUserTaskRemedyReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserTaskRemedyResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserTaskRemedy()

			_, _, err := moduleCli.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateUpdateAttendanceUserDailyShift(func(ctx context.Context, request *lark.CreateUpdateAttendanceUserDailyShiftReq, options ...lark.MethodOptionFunc) (*lark.CreateUpdateAttendanceUserDailyShiftResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateUpdateAttendanceUserDailyShift()

			_, _, err := moduleCli.CreateUpdateAttendanceUserDailyShift(ctx, &lark.CreateUpdateAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceGetAttendanceUserApproval(func(ctx context.Context, request *lark.GetAttendanceUserApprovalReq, options ...lark.MethodOptionFunc) (*lark.GetAttendanceUserApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceGetAttendanceUserApproval()

			_, _, err := moduleCli.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})

		t.Run("", func(t *testing.T) {
			cli.Mock().MockAttendanceCreateAttendanceUserApproval(func(ctx context.Context, request *lark.CreateAttendanceUserApprovalReq, options ...lark.MethodOptionFunc) (*lark.CreateAttendanceUserApprovalResp, *lark.Response, error) {
				return nil, nil, fmt.Errorf("mock-failed")
			})
			defer cli.Mock().UnMockAttendanceCreateAttendanceUserApproval()

			_, _, err := moduleCli.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.Equal(err.Error(), "mock-failed")
		})
	})

	t.Run("response is failed", func(t *testing.T) {
		cli := AppNoPermission.Ins()
		moduleCli := cli.Attendance

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateAttendanceUserSettings(ctx, &lark.UpdateAttendanceUserSettingsReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UploadAttendanceFile(ctx, &lark.UploadAttendanceFileReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateAttendanceGroup(ctx, &lark.CreateUpdateAttendanceGroupReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteAttendanceGroup(ctx, &lark.DeleteAttendanceGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceGroup(ctx, &lark.GetAttendanceGroupReq{
				GroupID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateAttendanceShift(ctx, &lark.CreateAttendanceShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.DeleteAttendanceShift(ctx, &lark.DeleteAttendanceShiftReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceShiftByID(ctx, &lark.GetAttendanceShiftByIDReq{
				ShiftID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceShiftByName(ctx, &lark.GetAttendanceShiftByNameReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceStatisticsData(ctx, &lark.GetAttendanceStatisticsDataReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceStatisticsHeader(ctx, &lark.GetAttendanceStatisticsHeaderReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.UpdateAttendanceUserStatisticsSettings(ctx, &lark.UpdateAttendanceUserStatisticsSettingsReq{
				UserStatsViewID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserStatisticsSettings(ctx, &lark.GetAttendanceUserStatisticsSettingsReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserDailyShift(ctx, &lark.GetAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserTask(ctx, &lark.GetAttendanceUserTaskReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserFlow(ctx, &lark.GetAttendanceUserFlowReq{
				UserFlowID: "x",
			})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchGetAttendanceUserFlow(ctx, &lark.BatchGetAttendanceUserFlowReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.BatchCreateAttendanceUserFlow(ctx, &lark.BatchCreateAttendanceUserFlowReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserTaskRemedy(ctx, &lark.GetAttendanceUserTaskRemedyReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateUpdateAttendanceUserDailyShift(ctx, &lark.CreateUpdateAttendanceUserDailyShiftReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.GetAttendanceUserApproval(ctx, &lark.GetAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})

		t.Run("", func(t *testing.T) {
			_, _, err := moduleCli.CreateAttendanceUserApproval(ctx, &lark.CreateAttendanceUserApprovalReq{})
			as.NotNil(err)
			as.True(lark.GetErrorCode(err) > 0, fmt.Sprintf("need get lark err, but get %s", err))
		})
	})
}
