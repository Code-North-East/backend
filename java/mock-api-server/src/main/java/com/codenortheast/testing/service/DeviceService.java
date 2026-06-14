package com.codenortheast.testing.service;


import com.codenortheast.testing.model.DeviceDTO;
import org.springframework.http.HttpStatus;
import org.springframework.stereotype.Service;
import org.springframework.web.server.ResponseStatusException;

import java.util.List;
import java.util.UUID;
import java.util.concurrent.CopyOnWriteArrayList;

@Service
public class DeviceService {

    // Thread-safe in-memory list acting as our mock database
    private final List<DeviceDTO> mockDatabase = new CopyOnWriteArrayList<>();

    // Pre-populate our mock database with seed data on initialization
    // Todo: change to a db later
    public DeviceService() {
        mockDatabase.add(new DeviceDTO("1", "MacBook Pro M3", 2024, null));
        mockDatabase.add(new DeviceDTO("2", "Google Pixel 8", 2023, null));
    }

    public List<DeviceDTO> getAllDevices() {
        return mockDatabase;
    }

    public DeviceDTO getDeviceById(String id) {
        return mockDatabase.stream()
                .filter(device -> device.getId().equals(id))
                .findFirst()
                .orElseThrow(() -> new ResponseStatusException(HttpStatus.NOT_FOUND, "Device not found with id: " + id));
    }

    public DeviceDTO createDevice(DeviceDTO deviceDTO) {
        // Automatically generate a safe random ID for testing
        // Change to a bson ID later
        deviceDTO.setId(UUID.randomUUID().toString());
        mockDatabase.add(deviceDTO);
        return deviceDTO;
    }
}
