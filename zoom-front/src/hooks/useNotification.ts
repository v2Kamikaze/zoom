import { notification } from "antd";

export type NotificationType = "success" | "info" | "warning" | "error";

export function useNotification() {
  const [api, notificationContext] = notification.useNotification();

  const openNotification = (
    type: NotificationType,
    title: string,
    message?: string
  ) => {
    api.destroy();

    api[type]({
      message: title,
      description: message,
      placement: "top",
    });
  };

  return {
    notificationContext,
    openNotification,
  };
}
