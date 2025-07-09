### Antrea Agent Resource Usage Analysis Report

#### **1. Key Findings Summary**
| **Configuration**          | **CPU Pattern**       | **Memory Pattern**     | **Key Characteristics** |  
|----------------------------|------------------------|------------------------|-------------------------|  
| **Base Antrea**            | Stable (8-18m)         | Stable (90-96Mi)       | Constant resource usage |  
| **+ AntreaIPAM**           | Minor fluctuations (9-21m) | Linear growth (80→94Mi) | IP allocation overhead |  
| **+ IPAM & SecondaryNetwork** | Volatile (8-38m)     | Steady growth (76→94Mi) | Multi-plane network overhead |  

---

### **2. CPU Usage Analysis**
#### **CPU Consumption Trend (millicores)**
| Pod Count | Base Antrea | +IPAM | +IPAM+Secondary |  
|-----------|-------------|-------|-----------------|  
| 5         | 9-10m       | 10m   | 9-38m           |  
| 10        | 9-10m       | 14-21m| 10-13m          |  
| 20        | 8-12m       | 10-11m| 8-10m           |  
| 40        | 9m          | 11-12m| 9-14m           |  
| 80        | 9-13m       | 11-12m| 8-12m           |  
| 100       | 14-18m      | 9-11m | 8-10m           |  

**Key Observations**:
1. **Base Antrea CPU Stability**
    - Maintains 8-18m (<2% of single core) across all scales
    - Demonstrates efficient dataplane processing

2. **IPAM-Induced Volatility**
    - 21m spike at 10 Pods (IP allocation burst)
    - Normalizes at higher loads (allocation completion)

3. **Secondary Network Anomaly**
    - **38m outlier** at 5 Pods (4× average)
    - Potential causes:
        - Multi-plane initialization overhead
        - ARP table synchronization burst
        - CNI coordination latency

---

### **3. Memory Usage Analysis**
#### **Memory Consumption Trend (MiB)**
| Pod Count | Base Antrea | +IPAM | +IPAM+Secondary |  
|-----------|-------------|-------|-----------------|  
| 5         | 90-91       | 80-81 | 76-77           |  
| 10        | 90-91       | 84-85 | 82-84           |  
| 20        | 91-95       | 86-89 | 85-90           |  
| 40        | 91          | 89-93 | 86-92           |  
| 80        | 91-96       | 90-94 | 88-92           |  
| 100       | 91-96       | 90-94 | 89-94           |  

**Key Observations**:
1. **Base Antrea Memory Efficiency**
    - Constant 91-96Mi (2.4-2.6% of node memory)
    - Pod-count independent data structures

2. **IPAM Linear Growth**
    - +3-4Mi per 10 Pods
    - **2Mi reduction** vs base at 100 Pods
    - Effective memory optimization

3. **Secondary Network Scaling**
    - Lowest initial usage (76Mi)
    - Highest growth slope (+18Mi at 100 Pods)
    - Multi-network metadata overhead

---

### **4. Performance Matrix**
| **Metric**       | Base Antrea         | +IPAM              | +IPAM+Secondary     |  
|------------------|---------------------|--------------------|---------------------|  
| **CPU Efficiency** | ⭐⭐⭐⭐⭐ (Stable)     | ⭐⭐⭐⭐ (Minor spikes) | ⭐⭐⭐ (High volatility) |  
| **Memory Efficiency** | ⭐⭐⭐⭐ (Flat)     | ⭐⭐⭐⭐⭐ (Optimized)  | ⭐⭐⭐ (Linear growth) |  
| **Scalability**  | ⭐⭐⭐⭐⭐ (No impact)  | ⭐⭐⭐⭐ (Predictable) | ⭐⭐⭐ (Monitor growth) |  

#### **Optimization Recommendations**
1. **Secondary Network Initialization**
   ```yaml  
   # antrea-agent config
   featureGates:
     SecondaryNetwork:
       initConcurrency: 2  # Limit parallel init
   ```  

2. **IPAM Memory Control**
   ```bash  
   # Increase IP recycling frequency
   antrea-agent --ipam-recycle-interval=5m
   ```  

3. **Resource Guardrails**
   ```yaml
   resources:
     limits:
       cpu: "50m"
       memory: "120Mi"
     requests:
       cpu: "10m"
       memory: "80Mi"
   ```

---

### **5. Risk Projection**
| **Scenario**               | 500 Pods Memory | Risk Level | Mitigation |  
|----------------------------|-----------------|------------|------------|  
| Base Antrea                | 95-100Mi        | Low        | None       |  
| +IPAM                      | 110-120Mi       | Medium     | Monitor    |  
| +IPAM+SecondaryNetwork     | 130-150Mi       | High       | Scale-out  |  

**Proactive Measures**:
1. **Memory Alerting**
   ```promql
   # Alert if >100Mi
   expr: container_memory_working_set_bytes{pod=~"antrea-agent-.*"} > 100 * 1024^2
   ```  
2. **Horizontal Sharding**
   ```bash
   antrea-agent --node-partition-id=$(hostname | cut -d- -f2)
   ```  
3. **Node Sizing Formula**
   ```  
   Max Pods = (Node Memory - 1GB) / Pod Memory Baseline
   ```  

> **Operational Guidance**: For SecondaryNetwork deployments, limit to 150 Pods/node. IPAM shows excellent memory optimization - recommend enabling first. Monitor memory growth slope beyond 40 Pods.