package com.codenortheast.testing.model;

import jakarta.validation.constraints.Max;
import jakarta.validation.constraints.Min;
import jakarta.validation.constraints.NotBlank;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor

public class DeviceDTO {
    private String id;
    @NotBlank(message = "Device title cannot be empty")
    private String title;
    @Max(value = 2026, message = "Invalid purchase year")
    @Min(value = 2010, message = "Invalid purchase year")
    private int purchaseYear;
    private Brand brand;
}
