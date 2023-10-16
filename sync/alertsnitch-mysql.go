package sync

/*
{
  "__inputs": [
    {
      "name": "DS_PROMETHEUS",
      "label": "Prometheus",
      "description": "",
      "type": "datasource",
      "pluginId": "prometheus",
      "pluginName": "Prometheus"
    }
  ],
  "__requires": [
    {
      "type": "panel",
      "id": "gauge",
      "name": "Gauge",
      "version": ""
    },
    {
      "type": "grafana",
      "id": "grafana",
      "name": "Grafana",
      "version": "6.1.3"
    },
    {
      "type": "panel",
      "id": "grafana-piechart-panel",
      "name": "Pie Chart",
      "version": "1.3.6"
    },
    {
      "type": "panel",
      "id": "graph",
      "name": "Graph",
      "version": ""
    },
    {
      "type": "datasource",
      "id": "prometheus",
      "name": "Prometheus",
      "version": "1.0.0"
    },
    {
      "type": "panel",
      "id": "singlestat",
      "name": "Singlestat",
      "version": ""
    }
  ],
  "annotations": {
    "list": [
      {
        "builtIn": 1,
        "datasource": "-- Grafana --",
        "enable": true,
        "hide": true,
        "iconColor": "rgba(0, 211, 255, 1)",
        "name": "Annotations & Alerts",
        "type": "dashboard"
      }
    ]
  },
  "editable": false,
  "gnetId": 10041,
  "graphTooltip": 0,
  "id": null,
  "iteration": 1555073446385,
  "links": [],
  "panels": [
    {
      "cacheTimeout": null,
      "colorBackground": false,
      "colorValue": false,
      "colors": [
        "#299c46",
        "rgba(237, 129, 40, 0.89)",
        "#d44a3a"
      ],
      "format": "none",
      "gauge": {
        "maxValue": 100,
        "minValue": 0,
        "show": false,
        "thresholdLabels": false,
        "thresholdMarkers": true
      },
      "gridPos": {
        "h": 7,
        "w": 4,
        "x": 0,
        "y": 0
      },
      "id": 21,
      "interval": null,
      "links": [],
      "mappingType": 1,
      "mappingTypes": [
        {
          "name": "value to text",
          "value": 1
        },
        {
          "name": "range to text",
          "value": 2
        }
      ],
      "maxDataPoints": 100,
      "nullPointMode": "connected",
      "nullText": null,
      "pluginVersion": "6.1.1",
      "postfix": "",
      "postfixFontSize": "50%",
      "prefix": "",
      "prefixFontSize": "50%",
      "rangeMaps": [
        {
          "from": "null",
          "text": "N/A",
          "to": "null"
        }
      ],
      "sparkline": {
        "fillColor": "rgba(31, 118, 189, 0.18)",
        "full": false,
        "lineColor": "rgb(31, 120, 193)",
        "show": false
      },
      "tableColumn": "",
      "targets": [
        {
          "expr": "gluster_volume_heal_count{volume='$volume',host='$hostname'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "thresholds": "",
      "timeFrom": null,
      "timeShift": null,
      "title": "gluster_volume_heal_count",
      "type": "singlestat",
      "valueFontSize": "80%",
      "valueMaps": [
        {
          "op": "=",
          "text": "N/A",
          "value": "null"
        }
      ],
      "valueName": "avg"
    },
    {
      "cacheTimeout": null,
      "gridPos": {
        "h": 7,
        "w": 20,
        "x": 4,
        "y": 0
      },
      "id": 17,
      "links": [],
      "options": {
        "maxValue": 100,
        "minValue": 0,
        "orientation": "auto",
        "showThresholdLabels": false,
        "showThresholdMarkers": true,
        "thresholds": [
          {
            "color": "green",
            "index": 0,
            "value": null
          },
          {
            "color": "dark-blue",
            "index": 1,
            "value": 25
          },
          {
            "color": "#EAB839",
            "index": 2,
            "value": 50
          },
          {
            "color": "dark-red",
            "index": 3,
            "value": 75
          }
        ],
        "valueMappings": [],
        "valueOptions": {
          "decimals": null,
          "prefix": "",
          "stat": "mean",
          "suffix": "",
          "unit": "percent"
        }
      },
      "pluginVersion": "6.1.3",
      "targets": [
        {
          "expr": "gluster_subvol_capacity_used_bytes{volume=\"$volume\",instance=\"$node\"} / gluster_subvol_capacity_total_bytes{volume=\"$volume\",instance=\"$node\"} * 100",
          "format": "table",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "gluster_subvol_capacity",
      "type": "gauge"
    },
    {
      "gridPos": {
        "h": 7,
        "w": 11,
        "x": 0,
        "y": 7
      },
      "id": 12,
      "links": [],
      "options": {
        "maxValue": 100,
        "minValue": 0,
        "orientation": "auto",
        "showThresholdLabels": true,
        "showThresholdMarkers": true,
        "thresholds": [
          {
            "color": "green",
            "index": 0,
            "value": null
          },
          {
            "color": "semi-dark-blue",
            "index": 1,
            "value": 25
          },
          {
            "color": "#EAB839",
            "index": 2,
            "value": 50
          },
          {
            "color": "dark-red",
            "index": 3,
            "value": 75
          }
        ],
        "valueMappings": [],
        "valueOptions": {
          "decimals": null,
          "prefix": "",
          "stat": "last",
          "suffix": "",
          "unit": "percent"
        }
      },
      "pluginVersion": "6.1.3",
      "targets": [
        {
          "expr": "gluster_cpu_percentage{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "gluster_cpu_percentage",
      "type": "gauge"
    },
    {
      "gridPos": {
        "h": 7,
        "w": 13,
        "x": 11,
        "y": 7
      },
      "id": 13,
      "links": [],
      "options": {
        "maxValue": 100,
        "minValue": 0,
        "orientation": "auto",
        "showThresholdLabels": true,
        "showThresholdMarkers": true,
        "thresholds": [
          {
            "color": "green",
            "index": 0,
            "value": null
          },
          {
            "color": "semi-dark-blue",
            "index": 1,
            "value": 25
          },
          {
            "color": "#EAB839",
            "index": 2,
            "value": 50
          },
          {
            "color": "dark-red",
            "index": 3,
            "value": 75
          }
        ],
        "valueMappings": [],
        "valueOptions": {
          "decimals": null,
          "prefix": "",
          "stat": "mean",
          "suffix": "",
          "unit": "percent"
        }
      },
      "pluginVersion": "6.1.3",
      "targets": [
        {
          "expr": "gluster_memory_percentage{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "gluster_memory_percentage",
      "type": "gauge"
    },
    {
      "aliasColors": {},
      "breakPoint": "50%",
      "cacheTimeout": null,
      "combine": {
        "label": "Others",
        "threshold": 0
      },
      "fontSize": "80%",
      "format": "bytes",
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 14
      },
      "id": 4,
      "interval": null,
      "legend": {
        "percentage": true,
        "show": true,
        "values": true
      },
      "legendType": "Right side",
      "links": [],
      "maxDataPoints": 3,
      "nullPointMode": "connected",
      "pieType": "pie",
      "strokeWidth": 1,
      "targets": [
        {
          "expr": "gluster_brick_inodes_free{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "legendFormat": "free",
          "refId": "A"
        },
        {
          "expr": "gluster_brick_inodes_used{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "legendFormat": "used",
          "refId": "B"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "gluster_brick_inodes",
      "transparent": true,
      "type": "grafana-piechart-panel",
      "valueName": "current"
    },
    {
      "aliasColors": {},
      "breakPoint": "50%",
      "cacheTimeout": null,
      "combine": {
        "label": "Others",
        "threshold": 0
      },
      "fontSize": "80%",
      "format": "bytes",
      "gridPos": {
        "h": 9,
        "w": 24,
        "x": 0,
        "y": 22
      },
      "id": 2,
      "interval": null,
      "legend": {
        "percentage": true,
        "show": true,
        "values": true
      },
      "legendType": "Right side",
      "links": [],
      "maxDataPoints": 3,
      "nullPointMode": "connected",
      "pieType": "pie",
      "pluginVersion": "6.1.1",
      "strokeWidth": 1,
      "targets": [
        {
          "expr": "gluster_brick_capacity_used_bytes{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "legendFormat": "used",
          "refId": "B"
        },
        {
          "expr": "gluster_brick_capacity_free_bytes{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "legendFormat": "free",
          "refId": "C"
        }
      ],
      "timeFrom": null,
      "timeShift": null,
      "title": "gluster_brick_capacity",
      "transparent": true,
      "type": "grafana-piechart-panel",
      "valueName": "current"
    },
    {
      "aliasColors": {},
      "bars": false,
      "cacheTimeout": null,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 31
      },
      "id": 19,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "gluster_virtual_memory_bytes{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "legendFormat": "",
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "gluster_virtual_memory_bytes",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "transparent": true,
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "bytes",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 39
      },
      "id": 15,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "gluster_resident_memory_bytes{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "gluster_resident_memory_bytes",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "bytes",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 8,
        "w": 24,
        "x": 0,
        "y": 47
      },
      "id": 10,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "gluster_elapsed_time_seconds{volume='$volume',instance='$node'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "gluster_elapsed_time_seconds",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    },
    {
      "aliasColors": {},
      "bars": false,
      "dashLength": 10,
      "dashes": false,
      "fill": 1,
      "gridPos": {
        "h": 9,
        "w": 24,
        "x": 0,
        "y": 55
      },
      "id": 6,
      "legend": {
        "avg": false,
        "current": false,
        "max": false,
        "min": false,
        "show": true,
        "total": false,
        "values": false
      },
      "lines": true,
      "linewidth": 1,
      "links": [],
      "nullPointMode": "null",
      "percentage": false,
      "pointradius": 2,
      "points": false,
      "renderer": "flot",
      "seriesOverrides": [],
      "spaceLength": 10,
      "stack": false,
      "steppedLine": false,
      "targets": [
        {
          "expr": "gluster_brick_up{volume='$volume'}",
          "format": "time_series",
          "interval": "$interval",
          "intervalFactor": 1,
          "refId": "A"
        }
      ],
      "thresholds": [],
      "timeFrom": null,
      "timeRegions": [],
      "timeShift": null,
      "title": "gluster_brick_up",
      "tooltip": {
        "shared": true,
        "sort": 0,
        "value_type": "individual"
      },
      "transparent": true,
      "type": "graph",
      "xaxis": {
        "buckets": null,
        "mode": "time",
        "name": null,
        "show": true,
        "values": []
      },
      "yaxes": [
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        },
        {
          "format": "short",
          "label": null,
          "logBase": 1,
          "max": null,
          "min": null,
          "show": true
        }
      ],
      "yaxis": {
        "align": false,
        "alignLevel": null
      }
    }
  ],
  "schemaVersion": 18,
  "style": "dark",
  "tags": [],
  "templating": {
    "list": [
      {
        "auto": false,
        "auto_count": 30,
        "auto_min": "10s",
        "current": {
          "text": "1m",
          "value": "1m"
        },
        "hide": 0,
        "label": "interval",
        "name": "interval",
        "options": [
          {
            "selected": true,
            "text": "1m",
            "value": "1m"
          },
          {
            "selected": false,
            "text": "10m",
            "value": "10m"
          },
          {
            "selected": false,
            "text": "30m",
            "value": "30m"
          },
          {
            "selected": false,
            "text": "1h",
            "value": "1h"
          },
          {
            "selected": false,
            "text": "6h",
            "value": "6h"
          },
          {
            "selected": false,
            "text": "12h",
            "value": "12h"
          },
          {
            "selected": false,
            "text": "1d",
            "value": "1d"
          },
          {
            "selected": false,
            "text": "7d",
            "value": "7d"
          },
          {
            "selected": false,
            "text": "14d",
            "value": "14d"
          },
          {
            "selected": false,
            "text": "30d",
            "value": "30d"
          }
        ],
        "query": "1m,10m,30m,1h,6h,12h,1d,7d,14d,30d",
        "refresh": 2,
        "skipUrlSync": false,
        "type": "interval"
      },
      {
        "allValue": null,
        "current": {},
        "datasource": "${DS_PROMETHEUS}",
        "definition": "label_values(gluster_brick_up,volume)",
        "hide": 0,
        "includeAll": false,
        "label": "数据卷",
        "multi": false,
        "name": "volume",
        "options": [],
        "query": "label_values(gluster_brick_up,volume)",
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "allValue": null,
        "current": {},
        "datasource": "${DS_PROMETHEUS}",
        "definition": "label_values(gluster_brick_up,hostname)",
        "hide": 0,
        "includeAll": false,
        "label": "主机",
        "multi": false,
        "name": "hostname",
        "options": [],
        "query": "label_values(gluster_brick_up,hostname)",
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      },
      {
        "allValue": null,
        "current": {},
        "datasource": "${DS_PROMETHEUS}",
        "definition": "label_values(up{job=\"glusterfs\"},instance)",
        "hide": 0,
        "includeAll": false,
        "label": "Node",
        "multi": false,
        "name": "node",
        "options": [],
        "query": "label_values(up{job=\"glusterfs\"},instance)",
        "refresh": 1,
        "regex": "",
        "skipUrlSync": false,
        "sort": 0,
        "tagValuesQuery": "",
        "tags": [],
        "tagsQuery": "",
        "type": "query",
        "useTags": false
      }
    ]
  },
  "time": {
    "from": "now-6h",
    "to": "now"
  },
  "timepicker": {
    "refresh_intervals": [
      "5s",
      "10s",
      "30s",
      "1m",
      "5m",
      "15m",
      "30m",
      "1h",
      "2h",
      "1d"
    ],
    "time_options": [
      "5m",
      "15m",
      "1h",
      "6h",
      "12h",
      "24h",
      "2d",
      "7d",
      "30d"
    ]
  },
  "timezone": "",
  "title": "glusterfs",
  "uid": "UaSar4RWk",
  "version": 2,
  "description": "Dashboard based on glusterfs official collector (https://github.com/gluster/gluster-prometheus)"
}*/

func TestMysql() {
}
