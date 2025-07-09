### **Test Report: Memory Impact of Enabling Antrea Features on Nodes**

---

#### **1. Test Overview**
| **Item**               | **Description**                                                                                  |  
|------------------------|--------------------------------------------------------------------------------------------------|  
| **Objective**          | Quantify memory impact of enabling `localPodInformer`, `AntreaIPAMController`, and `secondary network PodController` on nodes |  
| **Target Component**   | `antrea-agent` process memory usage                                                              |  
| **Features Tested**    | `localPodInformer`, `AntreaIPAMController`, `secondary network PodController`                      |  
| **Test Period**        | `[Start Date]` to `[End Date]`                                                                   |  
| **Test Owner**         | `[Name/Team]`                                                                                    |  

---

#### **2. Test Environment Configuration**
##### **2.1 Infrastructure**
```markdown  
- Kubernetes Version: __________  
- Antrea Version: __________  
- Node Specs: ____ vCPU, ____ GB RAM  
- Node Count: Control Plane: ____, Worker Nodes: ____  
- Network Plugin: __________  
```  

##### **2.2 Feature Configuration**
```yaml  
# antrea-agent ConfigMap Snippet  
antreaAgent:  
  extraArgs:  
    enable-local-pod-informer: "true"  

# antrea-controller ConfigMap Snippet  
antreaController:  
  extraArgs:  
    enable-antrea-ipam: "true"  
    enable-secondary-network: "true"  
```  

##### **2.3 Monitoring Tools**
```markdown  
- Prometheus Version: __________  
- Grafana Version: __________  
- Scrape Interval: ____ seconds  
```  

---

#### **3. Test Scenarios**
| **Scenario ID** | **Features Enabled**                                                                 | **Workload Description**                          | **Duration** |  
|-----------------|--------------------------------------------------------------------------------------|---------------------------------------------------|--------------|  
| Baseline        | None                                                                                 | No Pods                                           | 30 min       |  
| Scenario 1      | `localPodInformer` only                                                              | 50 standard Pods                                  | 60 min       |  
| Scenario 2      | `localPodInformer` + `AntreaIPAMController`                                          | 50 Pods using AntreaIPAM                          | 60 min       |  
| Scenario 3      | `localPodInformer` + `secondary network PodController`                               | 50 Pods with Secondary Networks                   | 60 min       |  
| Scenario 4      | All features                                                                         | Mixed workload (Standard + IPAM + Secondary Pods) | 24 hours     |  
| Stress Test     | All features                                                                         | 1000 Pods created/deleted at >50 Pods/min         | 120 min      |  

---

#### **4. Test Results**
##### **4.1 Baseline Memory Usage (MB)**
| **Scenario**   | **Idle** | **50 Pods** | **100 Pods** | **Δ (50 Pods)** |  
|----------------|----------|-------------|--------------|-----------------|  
| Baseline       | [Data]   | [Data]      | -            | [Data]%         |  
| Scenario 1     | [Data]   | [Data]      | [Data]       | [Data]%         |  
| Scenario 2     | [Data]   | [Data]      | [Data]       | [Data]%         |  
| Scenario 3     | [Data]   | [Data]      | [Data]       | [Data]%         |  
| Scenario 4     | [Data]   | [Data]      | [Data]       | [Data]%         |  

##### **4.2 Long-Running Memory Trend (Scenario 4)**
| **Time (h)**   | 0    | 6     | 12    | 18    | 24    |  
|----------------|------|-------|-------|-------|-------|  
| **Memory (MB)**| [Data] | [Data] | [Data] | [Data] | [Data] |  
| **Change (%)** | -    | [Data] | [Data] | [Data] | [Data] |  

##### **4.3 Stress Test Metrics**
| **Metric**                          | **Value**       |  
|-------------------------------------|-----------------|  
| Peak Pod Count                      | [Data]          |  
| Memory Peak                         | [Data] MB       |  
| Pod Churn Rate                      | [Data] Pods/min |  
| Memory Stabilization (after 50% Pod deletion) | [Data] min      |  

##### **4.4 Component Memory Breakdown (pprof)**
```markdown  
- localPodInformer cache: ____%  
- IPAM allocation records: ____%  
- NetworkController state: ____%  
- Goroutine overhead: ____%  
- Other: ____%  
```  

---

#### **5. Key Charts**
```mermaid  
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#FFAA60'}}}%%  
barChart  
    title Memory Comparison Across Scenarios (50 Pods)  
    x-axis Scenario  
    y-axis Memory (MB)  
    series Baseline: [Baseline Data]  
    series "Scenario 1": [Scenario 1 Data]  
    series "Scenario 2": [Scenario 2 Data]  
    series "Scenario 3": [Scenario 3 Data]  
    series "Scenario 4": [Scenario 4 Data]  
```  

```mermaid  
%%{init: {'theme': 'base', 'themeVariables': { 'primaryColor': '#70C3AF'}}}%%  
lineChart  
    title Scenario 4 Memory Trend (24h)  
    x-axis Time (h)  
    y-axis Memory (MB)  
    0: [0h Data]  
    6: [6h Data]  
    12: [12h Data]  
    18: [18h Data]  
    24: [24h Data]  
```  

---

#### **6. Findings**
```markdown  
1. [✓/✗] Memory Leak Detected:  
   - Location: __________  
   - Evidence: ____% growth over 24h (Threshold: 5%)  

2. [✓/✗] Stress Test Anomalies:  
   - Symptom: __________ (e.g., OOM kills, unreleased memory)  
   - Trigger Condition: Pod creation rate > ____/min  

3. [✓/✗] Feature Interaction Issues:  
   - Components: __________ + __________  
   - Impact: Additional ____% memory overhead  
```  

---

#### **7. Conclusions & Recommendations**
##### **7.1 Performance Impact**
```markdown  
- Feature Overhead:  
  • localPodInformer: +____ MB  
  • AntreaIPAMController: +____ MB  
  • SecondaryNetwork: +____ MB  
- Per-Pod Memory Cost: ____ MB/Pod  
- Recommended Node Sizing: ____ MB + (____ MB/Pod × Expected Pods)  
```  

##### **7.2 Optimization Recommendations**
```markdown  
High Priority:  
  • Fix memory leak in [Component] - Estimated savings: ____%  
  • Reduce informer cache TTL from [Current] to [Proposed]  

Medium Priority:  
  • Implement periodic cache pruning in __________  

General:  
  • Set antrea-agent memory limits:  
      resources:  
        limits:  
          memory: "____Mi"  
```  

##### **7.3 Risk Assessment**
```markdown  
Critical Risk:  
  • [Scenario] may exceed node memory at > ____ Pods  
  • Mitigation: Limit Secondary Network Pods to ____/node  

Safe Threshold:  
  • Max recommended memory utilization: ____% (____ MB buffer required)  
```  

---  

**Tested By**: __________________  
**Reviewed By**: __________________  
**Report Date**: `[YYYY-MM-DD]`

---  

### **How to Use This Template**
1. Replace all `[Data]` placeholders with measured values
2. Update Mermaid chart data series with actual results
3. Complete **Findings** section based on observed issues
4. Attach supplementary data:
    - Prometheus queries used
    - Heap profile snapshots (`pprof`)
    - Log excerpts showing critical events
    - Grafana dashboard exports

> **Note**: For long-term tests, record metrics hourly. For stress tests, capture peak values and recovery patterns.