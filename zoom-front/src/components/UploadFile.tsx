import React from "react";

export const UploadFile: React.FC<React.ComponentProps<"input">> = (props) => {
  return (
    <label htmlFor="file" className="py-4 hover:cursor-pointer">
      Upload
      <input
        id="file"
        type="file"
        className="hidden"
        accept="image/png, image/jpeg"
        {...props}
      />
    </label>
  );
};
